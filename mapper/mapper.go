package mapper

import (
	"context"
	"encoding/json"
)

func Map(ctx context.Context, input any, output any) error {
	rawBytes, err := json.Marshal(input)
	if err != nil {
		return err
	}

	return json.Unmarshal(rawBytes, output)
}
