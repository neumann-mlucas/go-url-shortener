package repository

import (
	"reflect"
	"testing"

	config "github.com/neumann-mlucas/go-url-shortener/internal/config"
	model "github.com/neumann-mlucas/go-url-shortener/internal/model"
	utils "github.com/neumann-mlucas/go-url-shortener/internal/utils"
)

var h1, h2, h3 string

func init() {
	h1, _ = utils.ToHash(1)
	h2, _ = utils.ToHash(2)
	h3, _ = utils.ToHash(3)
}

func Test_shortUrlRepository_CreateShortUrl(t *testing.T) {
	type args struct {
		fullurl string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{"add url", args{"https://gnu.org/"}, 1, false},
		{"duplicated url", args{"https://gnu.org/"}, 0, true},
		{"add another url", args{"https://linux.org/"}, 2, false},
	}

	// Initialize dependencies (clean in memory DB)
	config.LoadTestConfig()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &shortUrlRepository{
				db: config.AppConfig.DB,
			}
			got, err := r.CreateShortUrl(tt.args.fullurl)
			if (err != nil) != tt.wantErr {
				t.Errorf("shortUrlRepository.CreateShortUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("shortUrlRepository.CreateShortUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shortUrlRepository_GetShortUrlByID(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		want    *model.ShortUrl
		wantErr bool
	}{
		{"get ID 1", args{1}, &model.ShortUrl{Id: 1, Hash: h1, Url: "https://gnu.org/", Active: true}, false},
		{"get ID 2", args{2}, &model.ShortUrl{Id: 2, Hash: h2, Url: "https://linux.org/", Active: true}, false},
		{"get inactive URL", args{3}, nil, true},
		{"get nonexistence URL", args{4}, nil, true},
	}

	// Initialize dependencies (clean in memory DB)
	config.LoadTestConfig()
	insert_query := "INSERT INTO urls (hash, url, active) VALUES (?, ?, 1), (?, ?, 1), (?, ?, 0)"
	config.AppConfig.DB.Exec(insert_query, h1, "https://gnu.org/", h2, "https://linux.org/", h3, "https://neovim.io/")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &shortUrlRepository{
				db: config.AppConfig.DB,
			}
			got, err := r.GetShortUrlByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("shortUrlRepository.GetShortUrlByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortUrlRepository.GetShortUrlByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shortUrlRepository_GetShortUrlByHash(t *testing.T) {
	type args struct {
		hash string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.ShortUrl
		wantErr bool
	}{
		{"get hash gnu", args{h1}, &model.ShortUrl{Id: 1, Hash: h1, Url: "https://gnu.org/", Active: true}, false},
		{"get hash linux", args{h2}, &model.ShortUrl{Id: 2, Hash: h2, Url: "https://linux.org/", Active: true}, false},
		{"get inactive URL", args{h3}, nil, true},
		{"get nonexistence URL", args{"FOO="}, nil, true},
	}

	// Initialize dependencies (clean in memory DB)
	config.LoadTestConfig()
	insert_query := "INSERT INTO urls (hash, url, active) VALUES (?, ?, 1), (?, ?, 1), (?, ?, 0)"
	config.AppConfig.DB.Exec(insert_query, h1, "https://gnu.org/", h2, "https://linux.org/", h3, "https://neovim.io/")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &shortUrlRepository{
				db: config.AppConfig.DB,
			}
			got, err := r.GetShortUrlByHash(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("shortUrlRepository.GetShortUrlByHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortUrlRepository.GetShortUrlByHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shortUrlRepository_GetShortUrls(t *testing.T) {
	three_urls := []*model.ShortUrl{{Id: 1, Hash: h1, Url: "https://gnu.org/", Active: true}, {Id: 2, Hash: h2, Url: "https://linux.org/", Active: true}, {Id: 3, Hash: h3, Url: "https://neovim.io/", Active: false}}
	two_urls := []*model.ShortUrl{{Id: 1, Hash: h1, Url: "https://gnu.org/", Active: true}, {Id: 2, Hash: h2, Url: "https://linux.org/", Active: true}}

	type args struct {
		limit int
	}
	tests := []struct {
		name    string
		args    args
		want    []*model.ShortUrl
		wantErr bool
	}{
		{"get all urls", args{100}, three_urls, false},
		{"get with limit", args{2}, two_urls, false},
	}

	config.LoadTestConfig()
	insert_query := "INSERT INTO urls (hash, url, active) VALUES (?, ?, 1), (?, ?, 1), (?, ?, 0)"
	config.AppConfig.DB.Exec(insert_query, h1, "https://gnu.org/", h2, "https://linux.org/", h3, "https://neovim.io/")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &shortUrlRepository{
				db: config.AppConfig.DB,
			}
			got, err := r.GetShortUrls(tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("shortUrlRepository.GetShortUrls() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortUrlRepository.GetShortUrls() = %v, want %v", got, tt.want)
			}
		})
	}
}
