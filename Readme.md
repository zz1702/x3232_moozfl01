# x3232_moozfl01
Connect my linux PC ttyS to MAX3232CPE to RPi 2 B

moo is my main linux machine. It is on a Gigabyte (a Taiwan company I
think) motherboard GA-B85M-DS3H, which has a set of 9 pins of a COM
port, or ttyS0 on linux. As stated in the manual, Pin 2 is NSIN which
I think is RxD, and Pin 3 is NSOUT which I take as TxD. Pin 5 is
ground.

zfl01 is my RPi 2 B board. It has two pins (8 and 10) for TxD and
RxD. Pin 6 and Pin 9 are nearby ground pins. And Pin 1 provides 3.3V
supply from RPi.

I use a MAX3232CPE chip to make the +/- 12V on moo's COM port talk to
the 0~3.3V on RPi's TxD/RxD. The chip's Vcc uses RPi's 3.3V supply on
RPi's Pin 1. Notice that the ground on RPi and the ground on B85M
might be different, and then you must make sure that the difference
goes away before you connect the grounds together!

(zfl01 used an independent power supply and there's a ground
difference. So I have to make zfl01 get power from a USB hub connected
to moo. The hub is also independently supplied but there's almost no
difference in ground voltages of the hub and B85M's COM port.)

Then I tried minicom 2.6.1 on RPi, and minicom 2.7 on moo. And I also
tried to let linux kernel get hold of ttyS0 on RPi's side. In both
cases, moo can receive data but cannot send data. I also tried GNU
screen and played with setserial and stty settings. No
improvement. moo can receive but cannot send. Then I tried putty on
moo. Then both sending and receiving are O.K.! But putty gave me some
mosaic looking characters from time to time.

One major thing is, if I use the serial connection for login, then I
cannot use it for e.g. a PPP connection. If, on the other hand, I use
the line for transfering files, then I cannot use it for login. Login
using the serial line is visibly faster, more responsive, than login
through a wifi link. (You also don't need to get through e.g. an nmap
step to find out RPi's IP address.) Transfering large chunk of data
might not necessarily be faster given that the serial line only have
so much of a baud rate.

Anyway, that's the setup and situation. This project is for code I use
for moo and zfl01 to talk through the serial line.
