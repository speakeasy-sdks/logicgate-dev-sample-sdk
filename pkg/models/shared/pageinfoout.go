// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PageInfoOut - A collection page metadata
type PageInfoOut struct {
	// The zero-indexed page number (must not be less than 0, defaults to 0)
	Number *int64 `json:"number,omitempty"`
	// The size of the page and maximum number of items to be returned (must not be less than 1, defaults to 20)
	Size *int64 `json:"size,omitempty"`
	// The total number of items available
	TotalElements *int64 `json:"totalElements,omitempty"`
	// The total number of pages available based on the size
	TotalPages *int64 `json:"totalPages,omitempty"`
}

func (o *PageInfoOut) GetNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *PageInfoOut) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *PageInfoOut) GetTotalElements() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalElements
}

func (o *PageInfoOut) GetTotalPages() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalPages
}
