package service

import (
	"app/internal/core/model"
	"app/internal/core/port/outbound"
	errorMessage "app/internal/error"
	"errors"
	"math"
	"strings"

	"gonum.org/v1/gonum/mat"
)

/* Service for deciphering messages and calculating positions. */
type DecipherService struct {
	satelliteRepository outbound.SatellitePort
}

func NewDecipherService(repo outbound.SatellitePort) *DecipherService {
	return &DecipherService{
		satelliteRepository: repo,
	}
}

/*
GetLocation calculates the position of an object based on distance measurements from multiple satellites.
It uses the least squares method to solve a system of linear equations derived from the distance measurements.
The function returns the estimated position or an error if the calculation fails.
*/
func (ds *DecipherService) GetLocation(distances []model.Distance) (model.Position, error) {
	// Retrieve all satellites from the repository
	satellites := ds.satelliteRepository.GetAllSatellites()

	if len(satellites) == 0 {
		return model.Position{}, errors.New(errorMessage.PositionDetermined)
	}

	// Create matrices for linear system: A (coefficients) and b (right-hand side)
	A := mat.NewDense(len(satellites), 3, nil)
	b := mat.NewVecDense(len(satellites), nil)

	// Populate matrices A and b with data from the satellites and distances
	for i, satellite := range satellites {
		x, y := satellite.Position.X, satellite.Position.Y
		var d model.Distance
		found := false

		// Match each satellite with its corresponding distance measurement
		for _, dist := range distances {
			if strings.EqualFold(dist.Name, satellite.Name) {
				d = dist
				found = true
				break
			}
		}

		if found {
			// Populate matrix A with coefficients (2*x, 2*y, 1)
			// Populate vector b with the adjusted distance measurement
			A.Set(i, 0, 2*x)
			A.Set(i, 1, 2*y)
			A.Set(i, 2, 1)
			b.SetVec(i, x*x+y*y-d.Distance*d.Distance)
		}
	}
	// Create matrices for solving the linear system
	var ATA mat.Dense
	var ATb mat.VecDense
	var x mat.VecDense

	// Compute ATA = A^T * A and ATb = A^T * b
	ATA.Mul(A.T(), A)
	ATb.MulVec(A.T(), b)

	// Solve the system of equations ATA * x = ATb
	if err := x.SolveVec(&ATA, &ATb); err != nil {
		return model.Position{}, errors.New(errorMessage.PositionDetermined)
	}

	// Round the calculated position coordinates to two decimal places
	position := model.Position{
		X: math.Round(x.AtVec(0)*100) / 100,
		Y: math.Round(x.AtVec(1)*100) / 100,
	}

	// Return an error if the computed position is (0, 0)
	if position.X == 0 && position.Y == 0 {
		return model.Position{}, errors.New(errorMessage.PositionDetermined)
	}

	return position, nil
}

/*
GetMessage processes a slice of strings, where each inner slice represents a sequence of messages.
The function decodes the message ensuring that only unique, non-empty, and trimmed strings are included in the final message, avoiding consecutive duplicates.
It returns an error if no valid message can be determined.
*/
func (ds *DecipherService) GetMessage(messages [][]string) (string, error) {
	// Determine the maximum length of the inner slices
	maxLength := 0
	for _, msg := range messages {
		if len(msg) > maxLength {
			maxLength = len(msg)
		}
	}

	// Initialize an empty slice to store the decoded and unique trimmed messages
	var decodedMessage []string
	for i := 0; i < maxLength; i++ {
		for _, message := range messages {
			if i < len(message) {
				trimmed := strings.TrimSpace(message[i])
				if trimmed == "" {
					continue
				}

				decodedMessagelength := len(decodedMessage)

				// Get the last two elements of the decodedMessage slice, if they exist
				var prevDecodedMessage string
				if decodedMessagelength > 1 {
					prevDecodedMessage = decodedMessage[decodedMessagelength-2]
				}
				var lastDecodedMessage string
				if decodedMessagelength > 0 {
					lastDecodedMessage = decodedMessage[decodedMessagelength-1]
				}

				// Add the trimmed string to the decodeMessage slice if it's unique
				if decodedMessagelength == 0 || (prevDecodedMessage != trimmed && lastDecodedMessage != trimmed) {
					decodedMessage = append(decodedMessage, trimmed)
				}
			}
		}
	}

	// Return an error if no strings were added to the decodedMessage slice
	if len(decodedMessage) == 0 {
		return "", errors.New(errorMessage.MessageDetermined)
	}

	// Concatenate the decodedMessage strings with a space separator
	return strings.Join(decodedMessage, " "), nil
}

/*
GetSplitLocation retrieves the most recent messages from satellites, converts these messages into distance measurements, and calculates the object's location using these measurements.
It requires at least three messages to perform the calculation. If fewer than three messages are available, it returns an error.
*/
func (ds *DecipherService) GetSplitLocation() (model.Position, error) {
	// Retrieve the last messages received from satellites
	messages := ds.satelliteRepository.GetLastMessagesReceived()

	// Check if there are at least three messages to proceed with the location calculation
	if len(messages) < 3 {
		return model.Position{}, errors.New(errorMessage.LocationSplit)
	}

	// Convert the messages into a slice of model.Distance
	var distances []model.Distance
	for _, msg := range messages {
		distances = append(distances, model.Distance{Name: msg.Name, Distance: msg.Distance})
	}

	// Use the GetLocation method to calculate the position based on the distances
	return ds.GetLocation(distances)
}

func (ds *DecipherService) GetSplitMessage() (string, error) {
	// Retrieve the last messages received from satellites
	messages := ds.satelliteRepository.GetLastMessagesReceived()

	// Check if there are at least three messages to proceed with the message alignment
	if len(messages) < 3 {
		return "", errors.New(errorMessage.MessageSplit)
	}

	// Convert the messages into a slice of strings
	var messageArray [][]string
	for _, msg := range messages {
		messageArray = append(messageArray, msg.Message)
	}

	// Use the GetLocation method to align the messages
	return ds.GetMessage(messageArray)
}
