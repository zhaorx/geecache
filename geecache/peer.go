package geecache

type PeerPicker interface {
	PickPeer(key string) (peer peerGetter,ok bool)
}

type PeerGetter interface {
	Get(group string,key string) ([]byte,error)
}