package infra

import "cloud.google.com/go/firestore"

func (fs *Firestore) Set(userid string, c map[string]interface{}) error {
	if fs == nil {
		return ErrFirestore
	}
	_, err := fs.Client.Collection("users").Doc(userid).Set(fs.Context, c)
	return err
}

func (fs *Firestore) Update(userid string, path string, value interface{}) error {
	if fs == nil {
		return ErrFirestore
	}
	_, err := fs.Client.Collection("users").Doc(userid).Update(fs.Context, []firestore.Update{{
		Path:  path,
		Value: value,
	}})
	return err
}

func (fs *Firestore) Get(userid string) (map[string]interface{}, error) {
	dsnap, err := fs.Client.Collection("users").Doc(userid).Get(fs.Context)
	if err != nil {
		return nil, err
	}
	c := dsnap.Data()

	return c, nil
}

func (fs *Firestore) Close() {
	fs.Client.Close()
}
