package dbi

import (
	"os"
	"reflect"
	"testing"
)

func TestCreateURIFromEnv(t *testing.T) {
	user := "user"
	pw := "great"
	host := "localhost"
	port := "5432"
	database := "random"
	os.Setenv("PGUSER", user)
	os.Setenv("PGPASSWORD", pw)
	os.Setenv("PGHOST", host)
	os.Setenv("PGPORT", port)
	os.Setenv("PGDATABASE", database)

	expect := &URI{Username: user, Password: pw, Host: host, Port: port, Database: database}
	result := CreateURIFromEnv()

	if !reflect.DeepEqual(*expect, *result) {
		t.Error("Failed to create URI from env")
	}
}

func TestString(t *testing.T) {
	user := "user"
	pw := "great"
	host := "localhost"
	port := "5432"
	database := "random"

	expect1 := "postgresql://"
	expect2 := "postgresql://localhost"
	expect3 := "postgresql://localhost:5432"
	expect4 := "postgresql://localhost/random"
	expect5 := "postgresql://localhost:5432/random"
	expect6 := "postgresql://user@localhost"
	expect7 := "postgresql://user@localhost:5432"
	expect8 := "postgresql://user@localhost/random"
	expect9 := "postgresql://user@localhost:5432/random"
	expect10 := "postgresql://user:great@localhost"
	expect11 := "postgresql://user:great@localhost:5432"
	expect12 := "postgresql://user:great@localhost/random"
	expect13 := "postgresql://user:great@localhost:5432/random"

	r1 := &URI{Username: user}
	if r1.String() != expect1 {
		t.Error("Failed to stringify")
	}
	r2 := &URI{Password: pw}
	if r2.String() != expect1 {
		t.Error("Failed to stringify")
	}
	r3 := &URI{Host: host}
	if r3.String() != expect2 {
		t.Error("Failed to stringify")
	}
	r4 := &URI{Port: port}
	if r4.String() != expect1 {
		t.Error("Failed to stringify")
	}
	r5 := &URI{Database: database}
	if r5.String() != expect1 {
		t.Error("Failed to stringify")
	}
	r6 := &URI{Username: user, Password: pw}
	if r6.String() != expect1 {
		t.Error("Failed to stringify")
	}
	r7 := &URI{Username: user, Host: host}
	if r7.String() != expect6 {
		t.Error("Failed to stringify")
	}
	r8 := &URI{Username: user, Port: port}
	if r8.String() != expect1 {
		t.Error("Failed to stringify")
	}
	r9 := &URI{Username: user, Database: database}
	if r9.String() != expect1 {
		t.Error("Failed to stringify")
	}
	r10 := &URI{Password: pw, Host: host}
	if r10.String() != expect2 {
		t.Error("Failed to stringify")
	}
	r11 := &URI{Password: pw, Port: port}
	if r11.String() != expect1 {
		t.Error("Failed to stringify")
	}
	r12 := &URI{Password: pw, Database: database}
	if r12.String() != expect1 {
		t.Error("Failed to stringify")
	}
	r13 := &URI{Host: host, Port: port}
	if r13.String() != expect3 {
		t.Error("Failed to stringify")
	}
	r14 := &URI{Host: host, Database: database}
	if r14.String() != expect4 {
		t.Error("Failed to stringify")
	}
	r15 := &URI{Port: port, Database: database}
	if r15.String() != expect1 {
		t.Error("Failed to stringify")
	}
	r16 := &URI{Username: user, Password: pw, Host: host}
	if r16.String() != expect10 {
		t.Error("Failed to stringify")
	}
	r17 := &URI{Username: user, Password: pw, Port: port}
	if r17.String() != expect1 {
		t.Error("Failed to stringify")
	}
	r18 := &URI{Username: user, Password: pw, Database: database}
	if r18.String() != expect1 {
		t.Error("Failed to stringify")
	}
	r19 := &URI{Username: user, Host: host, Port: port}
	if r19.String() != expect7 {
		t.Error("Failed to stringify")
	}
	r20 := &URI{Username: user, Host: host, Database: database}
	if r20.String() != expect8 {
		t.Error("Failed to stringify")
	}
	r21 := &URI{Username: user, Port: port, Database: database}
	if r21.String() != expect1 {
		t.Error("Failed to stringify")
	}
	r22 := &URI{Password: pw, Host: host, Port: port}
	if r22.String() != expect3 {
		t.Error("Failed to stringify")
	}
	r23 := &URI{Password: pw, Host: host, Database: database}
	if r23.String() != expect4 {
		t.Error("Failed to stringify")
	}
	r24 := &URI{Password: pw, Port: port, Database: database}
	if r24.String() != expect1 {
		t.Error("Failed to stringify")
	}
	r25 := &URI{Host: host, Port: port, Database: database}
	if r25.String() != expect5 {
		t.Error("Failed to stringify")
	}
	r26 := &URI{Username: user, Password: pw, Host: host, Port: port}
	if r26.String() != expect11 {
		t.Error("Failed to stringify")
	}
	r27 := &URI{Username: user, Password: pw, Host: host, Database: database}
	if r27.String() != expect12 {
		t.Error("Failed to stringify")
	}
	r28 := &URI{Username: user, Password: pw, Port: port, Database: database}
	if r28.String() != expect1 {
		t.Error("Failed to stringify")
	}
	r29 := &URI{Username: user, Host: host, Port: port, Database: database}
	if r29.String() != expect9 {
		t.Error("Failed to stringify")
	}
	r30 := &URI{Password: pw, Host: host, Port: port, Database: database}
	if r30.String() != expect5 {
		t.Error("Failed to stringify")
	}
	r31 := &URI{Username: user, Password: pw, Host: host, Port: port, Database: database}
	if r31.String() != expect13 {
		t.Error("Failed to stringify")
	}
}
