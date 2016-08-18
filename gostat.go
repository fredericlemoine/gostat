/**
This package provides Statistical functions
*/
package gostats

import (
	crand "crypto/rand"
	"errors"
	"math"
	"math/big"
	rand "math/rand"
	"sort"
	// "strconv"
)

func Mean_int(a []int) float64 {
	var mean float64 = 0
	for _, v := range a {
		mean += float64(v)
	}
	return mean / float64(len(a))
}

func Mean_float(a []float64) float64 {
	var mean float64 = 0
	for _, v := range a {
		mean += v
	}
	return mean / float64(len(a))
}

func Median_int(a []int) float64 {
	/* we don't want to modify the original vector, so work on a copy that is going
	   to be sorted: */
	tmp := make([]int, len(a))
	copy(tmp, a)
	sort.Ints(tmp)
	return float64(tmp[int(math.Floor(float64(len(tmp))/2.0))]+tmp[int(math.Ceil(float64(len(tmp))/2.0))]) / 2.0
}

func Median_float(a []float64) float64 {
	/* we don't want to modify the original vector, so work on a copy that is going
	   to be sorted: */
	tmp := make([]float64, len(a))
	copy(tmp, a)
	sort.Float64s(tmp)
	return float64(tmp[int(math.Floor(float64(len(tmp))/2.0))]+tmp[int(math.Ceil(float64(len(tmp))/2.0))]) / 2.0
}

func Sum_int(a []int) int {
	var sum int = 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func Sum_float(a []float64) float64 {
	var sum float64 = 0
	for _, v := range a {
		sum += v
	}
	return sum
}

// Log factorial
func Log_fact(n int) float64 {
	var lf float64 = 0.0
	for i := 2; i <= n; i++ {
		lf = lf + math.Log(float64(i))
	}
	return lf
}

// Log factorial using Ramanujan approximation
func Factorial_log_rmnj(n int) float64 {
	if n == 0 {
		return 0.0
	} else if n <= 100 {
		return (Log_fact(n))
	} else {
		var accu float64 = 0.0
		accu += math.Log(float64(n)*(1.0+4.0*float64(n)*(1.0+2.0*float64(n)))+1.0/30.0-11.0/(240.0*float64(n))) / 6.0
		accu += math.Log(math.Pi) / 2.0
		accu -= float64(n)
		accu += float64(n) * math.Log(float64(n))
		return (accu)
	}
}

//
//  STAT FUNCTIONS
//

// Returns a random int in [0,n[
func RandTo(n int) (int64, error) {
	maxi := int64(math.MaxInt64)
	nBig, err := crand.Int(crand.Reader, big.NewInt(maxi))
	if err != nil {
		return -1, err
	}
	return nBig.Int64() % int64(n), nil
}

func Unif() float64 {
	var unif float64 = 0.5
	maxi := int64(math.MaxInt64)
	nBig, err := crand.Int(crand.Reader, big.NewInt(maxi))
	if err != nil {
		panic(err)
	}
	unif = (unif + (float64(nBig.Int64()))) / float64(maxi)
	return (unif)
}

func Exp(lambda float64) float64 {
	exp := Unif()
	exp = -math.Log(1-exp) / lambda
	return (exp)
}

func Gauss() float64 {
	unif1 := Unif()
	unif2 := Unif()
	gauss := math.Sqrt(-2*math.Log(unif1)) * math.Sin(2*math.Pi*(unif2))
	return (gauss)
}

func Normal(mu float64, sig float64) float64 {
	return (mu + (sig * Gauss()))
}

func Proba(p float64) bool {
	return (Unif() < p)
}

func Binomial(p float64, nb int) int {
	var binom int = 0
	for i := 0; i < nb; i++ {
		if Unif() < p {
			binom++
		}
	}
	return (binom)
}

// Samples num ints from the input of length length, with or without replacement
func Sample_int(data []int, num int, replace bool) ([]int, error) {
	if num < len(data) {
		return nil, errors.New("Sample Size must be >= data size")
	}

	output := make([]int, num)

	/* Without replacement */
	if !replace {
		// temp is a shuffled version of data
		perm := rand.Perm(len(data))
		for i := 0; i < num; i++ {
			output[i] = data[perm[i]]
		}
	} else {
		/* With replacement */
		for i := 0; i < num; i++ {
			r, _ := RandTo(len(data))
			output[i] = data[r]
		}
	}
	return output, nil
}

/* Shuffles the data in the array of length size */
func Shuffle_int(data []int) []int {
	temp := make([]int, len(data))
	perm := rand.Perm(len(data))
	for i, v := range perm {
		temp[v] = data[i]
	}
	return temp
}

func Shuffle_float(data []float64) []float64 {
	temp := make([]float64, len(data))
	perm := rand.Perm(len(data))
	for i, v := range perm {
		temp[v] = data[i]
	}
	return temp
}

func Sigma(values []float64) float64 {
	var vari float64 = 0.0
	mean := Mean_float(values)

	for _, v := range values {
		vari += math.Pow((v - mean), 2)
	}
	return (math.Sqrt(vari))
}

/* Original C++ implementation found at http://www.wilmott.com/messageview.cfm?catid=10&threadid=38771 */
/* C# implementation found at http://weblogs.asp.net/esanchez/archive/2010/07/29/a-quick-and-dirty-implementation-of-excel-norminv-function-in-c.aspx*/
/*
 *     Compute the quantile function for the normal distribution.
 *
 *     For small to moderate probabilities, algorithm referenced
 *     below is used to obtain an initial approximation which is
 *     polished with a final Newton step.
 *
 *     For very large arguments, an algorithm of Wichura is used.
 *
 *  REFERENCE
 *
 *     Beasley, J. D. and S. G. Springer (1977).
 *     Algorithm AS 111: The percentage points of the normal distribution,
 *     Applied Statistics, 26, 118-121.
 *
 *      Wichura, M.J. (1988).
 *      Algorithm AS 241: The Percentage Points of the Normal Distribution.
 *      Applied Statistics, 37, 477-484.
 */
/* Taken from https://gist.github.com/kmpm/1211922/ */
func Qnorm(p float64, mu float64, sigma float64) (float64, error) {
	var q, r, val float64

	if p < 0 || p > 1 {
		return -1, errors.New("Warning: p is < 0 or > 1 : returning DBL_MIN")
	}
	if sigma < 0 {
		return -1, errors.New("Warning: sigma is < 0 : returning NaN")
	}
	if p == 0 {
		return math.Inf(-1), nil
	}
	if p == 1 {
		return math.Inf(+1), nil
	}

	if sigma == 0 {
		return mu, nil
	}
	q = p - 0.5

	/*-- use AS 241 --- */
	/* double ppnd16_(double *p, long *ifault)*/
	/*      ALGORITHM AS241  APPL. STATIST. (1988) VOL. 37, NO. 3
	        Produces the normal deviate Z corresponding to a given lower
	        tail area of P; Z is accurate to about 1 part in 10**16.
	*/
	if math.Abs(q) <= .425 { /* 0.075 <= p <= 0.925 */
		r = .180625 - q*q
		val = q * (((((((r*2509.0809287301226727+
			33430.575583588128105)*r+67265.770927008700853)*r+
			45921.953931549871457)*r+13731.693765509461125)*r+
			1971.5909503065514427)*r+133.14166789178437745)*r + 3.387132872796366608) /
			(((((((r*5226.495278852854561+
				28729.085735721942674)*r+39307.89580009271061)*r+
				21213.794301586595867)*r+5394.1960214247511077)*r+
				687.1870074920579083)*r+42.313330701600911252)*r + 1)
	} else { /* closer than 0.075 from {0,1} boundary */
		/* r = min(p, 1-p) < 0.075 */
		if q > 0 {
			r = 1 - p
		} else {
			r = p
		}
		r = math.Sqrt(-math.Log(r))
		/* r = sqrt(-log(r))  <==>  min(p, 1-p) = exp( - r^2 ) */
		if r <= 5 { /* <==> min(p,1-p) >= exp(-25) ~= 1.3888e-11 */
			r += -1.6
			val = (((((((r*7.7454501427834140764e-4+
				.0227238449892691845833)*r+.24178072517745061177)*
				r+1.27045825245236838258)*r+
				3.64784832476320460504)*r+5.7694972214606914055)*
				r+4.6303378461565452959)*r +
				1.42343711074968357734) /
				(((((((r*1.05075007164441684324e-9+5.475938084995344946e-4)*r+
					.0151986665636164571966)*r+.14810397642748007459)*r+.68976733498510000455)*r+
					1.6763848301838038494)*r+2.05319162663775882187)*r + 1)
		} else { /* very close to  0 or 1 */
			r += -5
			val = (((((((r*2.01033439929228813265e-7+2.71155556874348757815e-5)*r+
				.0012426609473880784386)*r+.026532189526576123093)*r+.29656057182850489123)*r+
				1.7848265399172913358)*r+5.4637849111641143699)*r + 6.6579046435011037772) /
				(((((((r*2.04426310338993978564e-15+1.4215117583164458887e-7)*r+
					1.8463183175100546818e-5)*r+7.868691311456132591e-4)*r+.0148753612908506148525)*r+
					.13692988092273580531)*r+.59983220655588793769)*r + 1)
		}
		if q < 0.0 {
			val = -val
		}
	}
	return mu + sigma*val, nil
}

// From https://en.wikipedia.org/wiki/Normal_distribution
func Pnorm(x float64) float64 {
	var value, sum, result float64
	sum = x
	value = x
	for i := 1; i <= 100; i++ {
		value = (value * x * x / (2*float64(i) + 1))
		sum = sum + value
	}
	result = 0.5 + (sum/math.Sqrt(2*math.Pi))*math.Exp(-(x*x)/2)
	return result
}

// func main() {
// 	a := make([]int, 3)
// 	for i := 0; i < 3; i++ {
// 		a[i] = 30 - i
// 	}
// 	tmp := make([]int, len(a))
// 	copy(tmp, a)
// 	sort.Ints(tmp)

// 	for i := 0; i < 10000000; i++ {
// 		r, _ := RandTo(2000)
// 		fmt.Println(r)
// 	}
// 	// fmt.Println(Qnorm(0.002, 0, 1))
// 	// fmt.Println(Pnorm(1))

// 	// for i := 0; i < 1000; i++ {
// 	// 	fmt.Println(strconv.Itoa(i) + ": " + strconv.FormatFloat(Log_fact(i), 'f', 20, 64) + " | " + strconv.FormatFloat(Factorial_log_rmnj(i), 'f', 20, 64))
// 	// }
// }
