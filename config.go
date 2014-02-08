package wire

type Config struct {
	NumberOfInputters  int         `json:"number_of_inputters"`
	NumberOfOutputters int         `json:"number_of_outputters"`
	BufferSize         int         `json:"buffer_size"`
	In                 interface{} `json:"in"`
	Out                interface{} `json:"out"`
}

var DefaultConfig = Config{
	NumberOfInputters:  10,
	NumberOfOutputters: 10,
	BufferSize:         1024,
}
