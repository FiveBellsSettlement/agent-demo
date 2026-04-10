package demo

import "testing"

func TestTitleToken(t *testing.T) {
	t.Parallel()

	if got, want := TitleToken("demo-001"), "[task:demo-001]"; got != want {
		t.Fatalf("TitleToken() = %q, want %q", got, want)
	}
}

func TestPRTitle(t *testing.T) {
	t.Parallel()

	if got, want := PRTitle("demo-001", "add helper"), "[task:demo-001] add helper"; got != want {
		t.Fatalf("PRTitle() = %q, want %q", got, want)
	}
}

func TestBranchName(t *testing.T) {
	t.Parallel()

	if got, want := BranchName("demo-001"), "task/demo-001"; got != want {
		t.Fatalf("BranchName() = %q, want %q", got, want)
	}
}
