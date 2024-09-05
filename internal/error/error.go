package error

const (
	PositionDetermined       = "position cannot be determined"
	MessageDetermined        = "message cannot be determined"
	LocationSplit            = "not enough information to calculate the location"
	MessageSplit             = "not enough information to calculate the mesage"
	SatelliteNotFoundMessage = "message cannot be saved because satellite was not found"
	FirestoreRetrieving      = "error retrieving all documents from the collection"
	FirestoreSaving          = "error saving the document to the collection"
)
