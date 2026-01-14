package p2p
//Peer represents the
type Peer interface{

}

type Transport interface{
	ListenAndAccept() error
}