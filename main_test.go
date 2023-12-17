package main

import "testing"

func Test_getFirstdigit(t *testing.T) {
	type args struct {
		cadena string
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			name: "Base case",
			args: args{cadena: "1abc2"},
			want: '1',
		},	{
			name: "Base case",
			args: args{cadena: "pqr3stu8vwx"},
			want: '3',
		},	{
			name: "Base case",
			args: args{cadena: "a1b2c3d4e5f"},
			want: '1',
		},	{
			name: "Base case",
			args: args{cadena: "treb7uchet"},
			want: '7',
		},



	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFirstdigit(tt.args.cadena); got != tt.want {
				t.Errorf("getFirstdigit() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_getLastdigit(t *testing.T) {
	type args struct {
		cadena string
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			name: "Base case",
			args: args{cadena: "1abc2"},
			want: '2',
		},	{
			name: "Base case",
			args: args{cadena: "pqr3stu8vwx"},
			want: '8',
		},	{
			name: "Base case",
			args: args{cadena: "a1b2c3d4e5f"},
			want: '5',
		},	{
			name: "Base case",
			args: args{cadena: "treb7uchet"},
			want: '7',
		},



	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLastDigit(tt.args.cadena); got != tt.want {
				t.Errorf("getFirstdigit() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_reverse(t *testing.T){

  palindromo := "123"

  if palindromo == reverse(palindromo){
    t.Error("La funcion reverse nmo va")
  }


}
