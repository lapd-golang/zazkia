package main

import "io"

type clientAccess struct{}

func (c clientAccess) Read(link *link, r io.Reader, p parcel) (parcel, error) {
	if !link.transport.ReceivingFromClient {
		return emptyParcel, nil
	}
	return receiver{}.Read(link, r)
}

func (c clientAccess) Write(link *link, w io.Writer, p parcel) (parcel, error) {
	if p.isEmpty() {
		return p, nil
	}
	if !link.transport.SendingToClient {
		return emptyParcel, nil
	}
	return p, nil
}
