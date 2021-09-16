package main

import "testing"

var (
	inputIMUDPLog = []byte(`{ "name": "test_input_imudp", "origin": "imudp", "called.recvmmsg":1000, "called.recvmsg":2000, "msgs.received":500}`)
)

func TestgetInputIMUDP(t *testing.T) {
	logType := getStatType(inputIMUDPLog)
	if logType != rsyslogInputIMDUP {
		t.Errorf("detected pstat type should be %d but is %d", rsyslogInputIMDUP, logType)
	}

	pstat, err := newInputIMUDPFromJSON([]byte(inputLog))
	if err != nil {
		t.Fatalf("expected parsing input stat not to fail, got: %v", err)
	}

	if want, got := "test_input_imudp", pstat.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(1000), pstat.Recvmsg; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := int64(2000), pstat.Recvmmsg; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := int64(500), pstat.Received; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}
}

func TestInputIMUDPtoPoints(t *testing.T) {
	pstat, err := newInputIMUDPFromJSON([]byte(inputIMUDPLog))
	if err != nil {
		t.Fatalf("expected parsing input stat not to fail, got: %v", err)
	}

	points := pstat.toPoints()

	point := points[0]
	if want, got := "input_called_recvmmsg", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(1000), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := "test_input_imudp", point.LabelValue; want != got {
		t.Errorf("wanted '%s', got '%s'", want, got)
	}

	point = points[1]
	if want, got := "input_called_recvmsg", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(2000), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := "test_input_imudp", point.LabelValue; want != got {
		t.Errorf("wanted '%s', got '%s'", want, got)
	}

	point = points[2]
	if want, got := "input_received", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(500), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := "test_input_imudp", point.LabelValue; want != got {
		t.Errorf("wanted '%s', got '%s'", want, got)
	}
}
