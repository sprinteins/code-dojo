package datafieldmodel

// NewDatafieldModel _
func NewDatafieldModel(
	arkiID string,
	sspID string,
) DatafieldModel {

	return DatafieldModel{
		arkiID,
		sspID,
	}
}

// DatafieldModel Serves as template/prototype for the concrete Datafields
// zbps.: Market is a Datafield Model
type DatafieldModel struct {
	arkiID string
	sspID  string
}

// GetArkiID returns the foreign system id of Arki
func (dfm DatafieldModel) GetArkiID() string {
	return dfm.arkiID
}

// GetSspID returns the foreign system id of SSP
func (dfm DatafieldModel) GetSspID() string {
	return dfm.sspID
}
