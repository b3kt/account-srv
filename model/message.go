package model

// MsgPair - constant format
type MsgPair struct {
	code, msg interface{}
}

// MsgNoError - Message for no Error
var MsgNoError MsgPair = MsgPair{"00", "No Error"}

// MsgGeneralError - Message for general Error
var MsgGeneralError MsgPair = MsgPair{"01", "General Failure"}
