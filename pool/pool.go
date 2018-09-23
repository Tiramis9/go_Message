package pool
import (
    "errors"
    "log"
    "io"
    "sync"
)
type Pool struct {
    m sync.Mutex
    resources chan io.Closer
    factory func() (io.Closer,error)
    closed bool
}

var ErrPoolClosed = errors.New("Pool has been closed.")

func New(fn func() (io.Closer,error),size int) (*Pool,error) {
    if size <=0 {
        return nil,errors.New("size value too samall.")
    }
    return &Pool {
        factory : fn,
        resources: make (chan io.Closer,size),
    },nil
}

func(p *Pool)Acquire() (io.Closer,error) {
    select {
    case r,ok := <- p.resources:
        log.Println("Acquire:","Shared Resouce")
        if !ok {
            return nil,ErrPoolClosed
        }
        return r,nil
    
    default:
        log.Println("Acquire:","New Resource")
        return p.factory()
    }
}
func (p *Pool) Release(r io.Closer) {
    p.m.Lock()
    defer p.m.Unlock()

    if p.closed {
        r.Close()
        return
    }

    select {
    case p.resources <- r:
        log.Println("Release:","InQueue")

    default:
        log.Println("Release:","Closing")
        r.Close()
    }
}
func (p *Pool) Close() {
    p.m.Lock()
    defer p.m.Unlock()
    if p.closed {
        return
    }
    p.closed =true
    close(p.resources)
    for r:= range p.resources {
        r.Close()
    }
}

        


