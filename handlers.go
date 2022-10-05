package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	extraconnector "github.com/Nigel2392/extraconnector"
	"github.com/TwiN/go-color"
)

func handleSet(ask string) (extraconnector.Message, error) {
	set_syntax := strings.Split(ask, " ")
	if len(set_syntax) == 4 {
		key := set_syntax[1]
		value := set_syntax[2]
		ttl, err := strconv.Atoi(set_syntax[3])
		if err != nil {
			return extraconnector.Message{}, err
		}

		msg := extraconnector.Message{
			Channel_ID: SERVER.Current_Channel,
			Type:       "SET",
			Key:        key,
			Val:        value,
			TTL:        ttl,
		}
		return msg, nil
	} else {
		return extraconnector.Message{}, fmt.Errorf("invalid syntax")
	}
}

func handleGet(ask string) (extraconnector.Message, error) {
	get_syntax := strings.Split(ask, " ")
	if len(get_syntax) == 2 {
		key := get_syntax[1]
		msg := extraconnector.Message{
			Channel_ID: SERVER.Current_Channel,
			Type:       "GET",
			Key:        key,
		}
		return msg, nil
	} else {
		return extraconnector.Message{}, fmt.Errorf("invalid syntax")
	}
}

func handleDel(ask string) (extraconnector.Message, error) {
	del_syntax := strings.Split(ask, " ")
	if len(del_syntax) == 2 {
		key := del_syntax[1]
		msg := extraconnector.Message{
			Channel_ID: SERVER.Current_Channel,
			Type:       "DEL",
			Key:        key,
		}
		return msg, nil
	} else {
		return extraconnector.Message{}, fmt.Errorf("invalid syntax")
	}
}

func handleHasKey(ask string) (extraconnector.Message, error) {
	haskey_syntax := strings.Split(ask, " ")
	if len(haskey_syntax) == 2 {
		key := haskey_syntax[1]
		msg := extraconnector.Message{
			Channel_ID: SERVER.Current_Channel,
			Type:       "HASKEY",
			Key:        key,
		}
		return msg, nil
	} else {
		return extraconnector.Message{}, fmt.Errorf("invalid syntax")
	}
}
func handleSize(ask string) (extraconnector.Message, error) {
	size_syntax := strings.Split(ask, " ")
	if len(size_syntax) == 1 {
		msg := extraconnector.Message{
			Channel_ID: SERVER.Current_Channel,
			Type:       "SIZE",
		}
		return msg, nil
	} else {
		return extraconnector.Message{}, fmt.Errorf("invalid syntax")
	}

}
func handleSizeAll(ask string) (extraconnector.Message, error) {
	sizeall_syntax := strings.Split(ask, " ")
	if len(sizeall_syntax) == 1 {
		msg := extraconnector.Message{
			Channel_ID: SERVER.Current_Channel,
			Type:       "SIZEALL",
		}
		return msg, nil
	} else {
		return extraconnector.Message{}, fmt.Errorf("invalid syntax")
	}
}
func handleKeys(ask string) (extraconnector.Message, error) {
	keys_syntax := strings.Split(ask, " ")
	if len(keys_syntax) == 1 {
		msg := extraconnector.Message{
			Channel_ID: SERVER.Current_Channel,
			Type:       "KEYS",
		}
		return msg, nil
	} else {
		return extraconnector.Message{}, fmt.Errorf("invalid syntax")
	}
}
func handleSetChannel(ask string) error {
	setchannel_syntax := strings.Split(ask, " ")
	if len(setchannel_syntax) == 2 {
		channel, err := strconv.Atoi(setchannel_syntax[1])
		if err != nil {
			return errors.New("invalid syntax")
		}
		SERVER.Current_Channel = channel
		fmt.Println(color.Colorize(color.Purple, "Current channel changed to: "+strconv.Itoa(channel)))
		return nil
	}
	return errors.New("invalid syntax")
}
