package simulator

import "testing"

func TestRunOnceBadResults(t *testing.T) {
	vs := []struct {
		probabilityWin float64
		percentBet     float64
	}{
		{
			0.0,
			0.5,
		},
		{
			0.5,
			0.5,
		},
	}
	iterations := 1000

	for _, v := range vs {
		actual := RunOnce(v.probabilityWin, v.percentBet, iterations)
		if actual > 1 {
			t.Errorf("Playing bet: %f -- Probability to Win %f -- It should end in low results and it ended %f\n", v.percentBet, v.probabilityWin, actual)
		}
	}
}
func TestRunOnceGoodResults(t *testing.T) {
	vs := []struct {
		probabilityWin float64
		percentBet     float64
	}{
		{
			1.0,
			0.5,
		},
	}
	iterations := 1000

	for _, v := range vs {
		actual := RunOnce(v.probabilityWin, v.percentBet, iterations)
		if actual < 1 {
			t.Errorf("Playing bet: %f -- Probability to Win %f -- It should end in high results and it ended %f\n", v.percentBet, v.probabilityWin, actual)
		}
	}
}
