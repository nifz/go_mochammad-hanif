## Concurrent Programing

- Concurrency in Golang adalah tentang bagaimana membuat program berjalan secara bersamaan. Ada beberapa cara untuk menjalankan program yaitu secara sequential, parallel, dan concurrent. Namun hanya concurrent yang memungkinkan untuk dilakukan secara parallel.

- Salah satu fitur Go yang memungkinkan concurrent execution adalah goroutines. Goroutines adalah cara untuk menjalankan fungsi secara asynchronous yang memungkinkan program untuk terus berjalan sambil menyelesaikan tugas lainnya.

- Untuk melakukan komunikasi antar goroutines Go juga memiliki fitur channels dan select. Channels adalah cara untuk mengirim dan menerima data antar goroutines sedangkan select digunakan untuk memilih antara beberapa channel.

- Race condition adalah kondisi di mana dua atau lebih goroutines saling bersaing untuk mengakses suatu sumber daya atau data. Hal ini dapat menyebabkan hasil yang tidak konsisten dan tidak diinginkan. Untuk mengatasi masalah ini bisa digunakan teknik seperti blocking channels waitgroups dan mutex.

- Teknik untuk mengatasi race condition meliputi menghentikan channels, grup tunggu, dan penguncian data dengan menggunakan mutex.
