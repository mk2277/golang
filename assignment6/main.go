package main

import "fmt"

type person struct {
	firstname   string
	lastname    string
	fav_country []string
}

type vehicle struct {
	doors string
	color string
}

type truck struct {
	fourWheel     bool
	truck_vehicle vehicle
}

type sedan struct {
	luxury bool
	vehicle
}

type square struct {
	length  int
	breadth int
}

type circle struct {
	radius float64
}

func (x square) area() float64 {
	// fmt.Println(x.length * x.breadth)
	return float64(x.length) * float64(x.breadth)
}

func (x circle) area() float64 {
	// fmt.Println(3.14 * x.radius * x.radius)
	return 3.14 * x.radius * x.radius
}

type shape interface {
	area() float64
}

func info(i shape) float64 {
	// i.area()
	return i.area()
}

func main() {
	//1
	fmt.Println("First question")
	firstperson := person{
		firstname:   "John",
		lastname:    "Snow",
		fav_country: []string{"USA", "Australia", "Italy"},
	}

	secondperson := person{
		firstname:   "Tyrion",
		lastname:    "Lannister",
		fav_country: []string{"Austria", "Turkey", "Dubai"},
	}
	fmt.Println("first person struct", firstperson)
	fmt.Println("Person First Name ", firstperson.firstname)
	fmt.Println("Lopping over the favourite Country")
	for i := range firstperson.fav_country {

		fmt.Println(i, firstperson.fav_country[i])

	}
	fmt.Println("second person struct", secondperson)
	fmt.Println("Person First Name ", secondperson.firstname)
	fmt.Println(secondperson.lastname)
	fmt.Println("Lopping over the favourite Country")
	for i := range secondperson.fav_country {

		fmt.Println(i, secondperson.fav_country[i])

	}
	//2
	fmt.Println(" ")
	fmt.Println("Second question")

	a := make(map[string]person)

	a[firstperson.lastname] = firstperson

	a[secondperson.lastname] = secondperson

	fmt.Println(a)

	for i, j := range a {
		// fmt.Println(a[i].firstname, a[i].lastname)
		fmt.Println("person struct", a[i])
		fmt.Println("person first name ", a[i].firstname)
		fmt.Println("looping over the favourite country")
		for x := range j.fav_country {
			fmt.Println(x, j.fav_country[x])
		}
	}

	//3
	fmt.Println("")
	fmt.Println("Third question")
	chevy := truck{
		fourWheel:     true,
		truck_vehicle: vehicle{"fourdoor", "red"},
	}

	toyota := sedan{
		luxury:  true,
		vehicle: vehicle{"twodoor", "black"},
	}
	fmt.Println(toyota)
	fmt.Println(chevy, toyota)
	fmt.Println(chevy.fourWheel)
	fmt.Println(chevy.truck_vehicle)
	fmt.Println(chevy.truck_vehicle.doors)

	//4
	fmt.Println("")
	fmt.Println("fourth question")

	sqrvalues := square{
		length:  10,
		breadth: 10,
	}

	circvalues := circle{
		radius: 10,
	}

	// fmt.Println("area of square", sqrvalues.area())
	// fmt.Println("area of circle", circvalues.area())
	fmt.Println(info(sqrvalues))
	fmt.Println(info(circvalues))

	// p := info(sqrvalues)
	// info(sqrvalues)
	// info(circvalues)

}
