// Create a combat function that takes the player's current health and the amount of damage recieved, and returns the player's new health. Health can't be less than 0.

package kata

func combat(health, damage float64) float64 {
	healthPost := health - damage
	if healthPost < 0 {
		return float64(0)
	}
	return float64(healthPost)
}
