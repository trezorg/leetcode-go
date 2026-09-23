// Идея решения: двигать по s окно длиной len(p). Если в окне каждая буква встречается столько же раз, сколько в p, значит это анаграмма.
//
//	subStringCounts хранит частоты букв в p, а stringCounts — в текущем окне s. Индекс 0 соответствует a, индекс 1 — b и так далее. Сначала код заполняет оба массива для первого окна и, если они
//	равны, добавляет индекс 0.
//
//	Затем окно сдвигается на один символ вправо. Для нового начала i код убирает из подсчёта ушедший символ s[i-1], добавляет вошедший s[i+n-1] и снова сравнивает массивы. При равенстве i попадает
//	в ответ. Так не приходится пересчитывать частоты всего окна после каждого сдвига.
//
//	Например, для s = "cbaebabacd" и p = "abc" совпадут окна "cba" и "bac" — результат [0, 6].
//
//	Массивы имеют фиксированный размер 26, поэтому время работы — O(len(p) + 26·len(s)), а дополнительная память без учёта ответа — O(26). Решение предполагает, что строки содержат строчные
//	латинские буквы.
package main

const size = 'z' - 'a' + 1

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	n, sn := len(p), len(s)
	if sn < n {
		return res
	}
	subStringCounts := [size]int{}
	stringCounts := [size]int{}

	for i := 0; i < n; i++ {
		subStringCounts[p[i]-'a']++
		stringCounts[s[i]-'a']++
	}

	if subStringCounts == stringCounts {
		res = append(res, 0)
	}

	for i := 1; i < sn-n+1; i++ {
		stringCounts[s[i-1]-'a']--
		stringCounts[s[i+n-1]-'a']++
		if subStringCounts == stringCounts {
			res = append(res, i)
		}
	}

	return res
}
