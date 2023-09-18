package protobuf

import (
	"fmt"
	"testing"

	. "github.com/easierway/partialparse/protobuf_def"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PrepareTestData() ([]byte, error) {
	now := timestamppb.Now()
	p := Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*Person_PhoneNumber{
			{
				Number: "555-4321",
				Type:   Person_PHONE_TYPE_HOME,
			},
			{
				Number: "555-4321",
				Type:   Person_PHONE_TYPE_HOME,
			},
			{
				Number: "555-4321",
				Type:   Person_PHONE_TYPE_HOME,
			},
			{
				Number: "555-4321",
				Type:   Person_PHONE_TYPE_HOME,
			},
			{
				Number: "555-4321",
				Type:   Person_PHONE_TYPE_HOME,
			},
			{
				Number: "555-4321",
				Type:   Person_PHONE_TYPE_HOME,
			},
		},
		LastUpdated: now,
	}
	return proto.Marshal(&p)
}

func TestFullyParse(t *testing.T) {
	data, err := PrepareTestData()
	if err != nil {
		t.Fatal(err)
	}
	p1 := &Person{}
	proto.Unmarshal(data, p1)
	fmt.Println(p1)
}

func TestPartialParse(t *testing.T) {
	data, err := PrepareTestData()
	if err != nil {
		t.Fatal(err)
	}
	p2 := &PartialPerson{}
	proto.Unmarshal(data, p2)
	fmt.Println(p2)
}

func BenchmarkFullyParse(b *testing.B) {
	data, err := PrepareTestData()
	if err != nil {
		b.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		p1 := &Person{}
		proto.Unmarshal(data, p1)
	}
}

func BenchmarkPartialParse(b *testing.B) {
	data, err := PrepareTestData()

	if err != nil {
		b.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		p2 := &PartialPerson{}
		proto.Unmarshal(data, p2)
	}
}
