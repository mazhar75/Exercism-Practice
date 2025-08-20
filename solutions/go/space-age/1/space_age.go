package space

type Planet string
const sec float64=31557600
func GiveYearsRelation() map[string]float64{
    years := map[string]float64{
        "Mercury": 0.2408467,
        "Venus":   0.61519726,
        "Earth":   1.0,
        "Mars":    1.8808158,
        "Jupiter": 11.862615,
        "Saturn":  29.447498,
        "Uranus":  84.016846,
        "Neptune": 164.79132,
    }
    return years
}

func Age(seconds float64, planet Planet) float64 {
	relation := GiveYearsRelation()
    earthYears := seconds / sec
    val, ok := relation[string(planet)]
    if ok{
         planetYears := earthYears / val
         return planetYears
    } else {
        return -1
    }
    
   
}
