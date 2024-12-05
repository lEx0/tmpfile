package tmpfile

import (
	"io"
	"strings"
	"testing"
)

func TestNewFromReader(t *testing.T) {
	f, err := NewFromReader(strings.NewReader("test"))
	if err != nil {
		t.Fatalf("tmpfile must be created: %v", err)
	}

	body, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("tmpfile must be readable: %v", err)
	} else if string(body) != "test" {
		t.Fatalf("tmpfile must contain 'test' but got %s", string(body))
	}

	_, err = f.file.Stat()
	if err != nil {
		t.Fatalf("tmpfile must exist: %v", err)
	}

	if err := f.Close(); err != nil {
		t.Fatalf("tmpfile must be closable: %v", err)
	}

	// проверяем что файл удалился
	_, err = f.file.Stat()
	if err == nil {
		t.Fatalf("tmpfile must be deleted")
	}
}

func TestNew(t *testing.T) {
	f, err := New()
	if err != nil {
		t.Fatalf("tmpfile must be created: %v", err)
	}

	_, err = io.Copy(f, strings.NewReader("test"))
	if err != nil {
		t.Fatalf("tmpfile must be writable: %v", err)
	}

	// читаем из файла
	body, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("tmpfile must be readable: %v", err)
	} else if string(body) != "test" {
		t.Fatalf("tmpfile must contain 'test' but got %s", string(body))
	}

	_, _ = f.Seek(0, io.SeekStart)
	body, err = io.ReadAll(f)
	if err != nil {
		t.Fatalf("tmpfile must be readable: %v", err)
	} else if string(body) != "test" {
		t.Fatalf("tmpfile must contain 'test' but got %s", string(body))
	}

	_, err = f.file.Stat()
	if err != nil {
		t.Fatalf("tmpfile must exist: %v", err)
	}

	if err = f.Close(); err != nil {
		t.Fatalf("tmpfile must be closable: %v", err)
	}

	// проверяем что файл удалился
	_, err = f.file.Stat()
	if err == nil {
		t.Fatalf("tmpfile must be deleted")
	}
}

func TestSeek(t *testing.T) {
	f, err := New()
	if err != nil {
		t.Fatalf("tmpfile must be created: %v", err)
	}

	_, err = io.Copy(f, strings.NewReader("hello world!"))
	if err != nil {
		t.Fatalf("tmpfile must be writable: %v", err)
	}
	buf := make([]byte, 6)

	// проверяем чтение с середины файла
	_, err = f.Seek(6, io.SeekStart)
	if err != nil {
		t.Fatalf("tmpfile must be seekable: %v", err)
	}

	readCount, err := f.Read(buf)
	if err != nil {
		t.Fatalf("tmpfile must be readable: %v", err)
	} else if readCount != 6 {
		t.Fatalf("tmpfile must read 6 bytes but got %d", readCount)
	} else if string(buf) != "world!" {
		t.Fatalf("tmpfile must contain 'world!' but got %s", string(buf))
	}

	buf = make([]byte, 1)
	_, err = f.Seek(-1, io.SeekEnd)
	if err != nil {
		t.Fatalf("tmpfile must be seekable: %v", err)
	} else if _, err = f.Read(buf); err != nil {
		t.Fatalf("tmpfile must be readable: %v", err)
	} else if string(buf) != "!" {
		t.Fatalf("tmpfile must contain '!' but got %s", string(buf))
	}

	// проверяем чтение файла целиком
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		t.Fatalf("tmpfile must be seekable: %v", err)
	} else if body, err := io.ReadAll(f); err != nil {
		t.Fatalf("tmpfile must be readable: %v", err)
	} else if string(body) != "hello world!" {
		t.Fatalf("tmpfile must contain 'hello world!' but got %s", string(body))
	}
}
