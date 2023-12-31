// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PageLinksOut - A collection of page URL links for navigation and iteration
type PageLinksOut struct {
	// A URL path to the first page of requested data
	First *string `json:"first,omitempty"`
	// A URL link to the last page of requested data
	Last *string `json:"last,omitempty"`
	// A URL link to the next page of requested data or `null` if currently on last page
	Next *string `json:"next,omitempty"`
	// A URL path to the next page of requested data or `null` if currently on first page
	Prev *string `json:"prev,omitempty"`
	// A URL path to the current page of requested data
	Self *string `json:"self,omitempty"`
}

func (o *PageLinksOut) GetFirst() *string {
	if o == nil {
		return nil
	}
	return o.First
}

func (o *PageLinksOut) GetLast() *string {
	if o == nil {
		return nil
	}
	return o.Last
}

func (o *PageLinksOut) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *PageLinksOut) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

func (o *PageLinksOut) GetSelf() *string {
	if o == nil {
		return nil
	}
	return o.Self
}
