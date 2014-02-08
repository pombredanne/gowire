package wire

type Wire struct {
	in  *input
	out *output

	Config *Config
}

func New(config *Config) *Wire {
	w := new(Wire)
	if config == nil {
		w.Config = &DefaultConfig
	} else {
		w.Config = config
	}
	return w
}

func (w *Wire) Start(in Inputter, out Outputter, transformer Transformer, errs chan error) error {
	messages := make(chan interface{}, w.Config.BufferSize)

	w.in = newInput(in, transformer)
	if err := w.in.start(w.Config.NumberOfInputters, messages, errs); err != nil {
		return err
	}

	w.out = newOutput(out, in)
	if err := w.out.start(w.Config.NumberOfOutputters, messages, errs); err != nil {
		return err
	}

	return nil
}

func (w *Wire) Close() error {
	if err := w.in.close(); err != nil {
		return err
	}
	return w.out.close()
}
