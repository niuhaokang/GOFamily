package TN

import (
	"fmt"
	"testing"
)

var (
	x = []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 3443, 3435, 32322, 5543, 23345, 3232, 55423, 23324,
	}
)

func TestTOPKQuickSelect(t *testing.T) {
	fmt.Println(TOPKQuickSelect(x, 30))
}

//  8472 ns/op
func BenchmarkTOPKQuickSelect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TOPKQuickSelect(x, 30)
	}
}
