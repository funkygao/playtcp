TCP Handshake
=============


### Why handshake?

相互同步起始seqnum和recv window size


### Why not 2-way handshake?

    A -> B

    如果B返回的 syn+ack没有到达A，那么：
    B认为该连接已经建立完成，而A认为连接没有完成。
    此时，B可以开始发送PUSH，而A一直等待SYN+ACK到达，不会接受PUSH，造成死锁


### Why seqnum randomized instead of from 0/1?

If host crash and restart, it will confuse remote endpoint into believing that
old connection remained open.

Besides, a malicious person could write code to analyze ISNs and then predict the ISN of a 
subsequent TCP connection based on the ISNs used in earlier ones. 
