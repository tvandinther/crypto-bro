package main

func (a *analyser) processText(text string) ([]string, int, error) {
	crypto, err := a.IdentifyCrypto(text)
	if err != nil {
		return nil, 0, err
	}

	sentiment := a.CalculateSentiment(text)

	return crypto, sentiment, nil
}