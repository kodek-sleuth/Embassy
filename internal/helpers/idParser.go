package helpers

import uuid "github.com/satori/go.uuid"

func ParseIDs(ids []string) ([]uuid.UUID, error){
	var parsedIDS []uuid.UUID
	for _, v := range ids {
		parsedID, err := uuid.FromString(v)
		if err != nil{
			return nil, err
		}
		parsedIDS = append(parsedIDS, parsedID)
	}
	return parsedIDS, nil
}
