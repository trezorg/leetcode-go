package main

func canBeValid(s string, locked string) bool {
	if len(s)%2 != 0 {
		return false
	}
	balanced_forward, free_forward, ss_forward := 0, 0, byte('(')
	balanced_backward, free_backward, ss_backward := 0, 0, byte(')')
	for i, j := 0, len(s)-1; i < len(s) && j >= 0; i, j = i+1, j-1 {
		if locked[i] == '1' {
			if s[i] == ss_forward {
				balanced_forward++
			} else {
				balanced_forward--
			}
		} else {
			free_forward++
		}
		if free_forward+balanced_forward < 0 {
			return false
		}
		if locked[j] == '1' {
			if s[j] == ss_backward {
				balanced_backward++
			} else {
				balanced_backward--
			}
		} else {
			free_backward++
		}
		if free_backward+balanced_backward < 0 {
			return false
		}
	}
	if balanced_forward > free_forward || balanced_backward > free_backward {
		return false
	}
	return true
}
