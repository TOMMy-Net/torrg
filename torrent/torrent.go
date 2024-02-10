package torrent

import (
	"bytes"
	"crypto/sha1"
	"fmt"

	"os"

	bencode "github.com/jackpal/bencode-go"
)

type bencodeInfo struct {
	Pieces      string `bencode:"pieces"`
	PieceLength int    `bencode:"piece length"`
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
}

type bencodeFile struct {
	Announce string      `bencode:"announce"`
	Info     bencodeInfo `bencode:"info"`
}

type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

func Open(path string) TorrentFile {
	file, err := os.Open(path)

	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	BFile := bencodeFile{}
	err = bencode.Unmarshal(file, &BFile)
	if err != nil {
		os.Exit(1)
	}
	f, err := BFile.toTorrentFile()
	if err != nil {
		return TorrentFile{}
	}
	return f
}

func (I *bencodeInfo) splitPieceHash() ([][20]byte, error) {
	hashLen := 20
	buf := []byte(I.Pieces)

	if len(buf)%hashLen != 0 {
		return nil, fmt.Errorf("Invalid pieces length: %d", len(buf))
	}
	numHashes := len(buf) / hashLen
	hashes := make([][20]byte, numHashes)
	for i := 0; i < numHashes; i++ {
		copy(hashes[i][:], buf[i*hashLen:(i+1)*hashLen])
	}
	return hashes, nil
}

func (I *bencodeInfo) hashInfo() ([20]byte, error) {
	var buf bytes.Buffer
	err := bencode.Marshal(&buf, I)
	if err != nil {
		return [20]byte{}, err
	}
	var h = sha1.Sum(buf.Bytes())
	return h, nil
}

func (BFile *bencodeFile) toTorrentFile() (TorrentFile, error) {
	info, err := (BFile.Info).hashInfo()
	if err != nil {
		return TorrentFile{}, err
	}

	pieceHash, err := (BFile.Info).splitPieceHash()
	if err != nil {
		return TorrentFile{}, err
	}

	t := TorrentFile{
		Announce:    BFile.Announce,
		InfoHash:    info,
		PieceHashes: pieceHash,
		PieceLength: BFile.Info.PieceLength,
		Length:      BFile.Info.Length,
		Name:        BFile.Info.Name,
	}
	return t, nil
}
