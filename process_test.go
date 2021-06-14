package main

import (
	"strings"
	"testing"
	"time"
)

func TestGetProcess(t *testing.T) {
	out := `      1 root      16  -4  171220  11884   6164 S   0.0   0.3   0:16.91 systemd
      2 root      20   0       0      0      0 S   0.0   0.0   0:00.03 kthreadd
      3 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 rcu_gp
      4 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 rcu_par_gp
      6 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 kworker/0:0H-kblockd
      9 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 mm_percpu_wq
     10 root      20   0       0      0      0 S   0.0   0.0   0:11.76 ksoftirqd/0
     11 root      20   0       0      0      0 I   0.0   0.0   1:55.50 rcu_sched
     12 root      rt   0       0      0      0 S   0.0   0.0   0:00.92 migration/0
     13 root     -51   0       0      0      0 S   0.0   0.0   0:00.00 idle_inject/0
     14 root      20   0       0      0      0 S   0.0   0.0   0:00.00 cpuhp/0
     15 root      20   0       0      0      0 S   0.0   0.0   0:00.00 cpuhp/1
     16 root     -51   0       0      0      0 S   0.0   0.0   0:00.00 idle_inject/1
     17 root      rt   0       0      0      0 S   0.0   0.0   0:01.02 migration/1
     18 root      20   0       0      0      0 S   0.0   0.0   0:13.05 ksoftirqd/1
     20 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 kworker/1:0H-kblockd
     21 root      20   0       0      0      0 S   0.0   0.0   0:00.00 kdevtmpfs
     22 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 netns
     23 root      20   0       0      0      0 S   0.0   0.0   0:00.00 rcu_tasks_kthre
     24 root      20   0       0      0      0 S   0.0   0.0   0:00.00 kauditd
     25 root      20   0       0      0      0 S   0.0   0.0   0:00.10 khungtaskd
     26 root      20   0       0      0      0 S   0.0   0.0   0:00.68 oom_reaper
     27 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 writeback
     28 root      20   0       0      0      0 S   0.0   0.0   0:00.00 kcompactd0
     29 root      25   5       0      0      0 S   0.0   0.0   0:00.00 ksmd
     30 root      39  19       0      0      0 S   0.0   0.0   0:00.91 khugepaged
     77 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 kintegrityd
     78 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 kblockd
     79 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 blkcg_punt_bio
     80 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 tpm_dev_wq
     81 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 ata_sff
     82 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 md
     83 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 edac-poller
     84 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 devfreq_wq
     85 root      rt   0       0      0      0 S   0.0   0.0   0:00.00 watchdogd
     88 root      20   0       0      0      0 S   0.0   0.0   3:50.73 kswapd0
     89 root      20   0       0      0      0 S   0.0   0.0   0:00.00 ecryptfs-kthrea
     91 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 kthrotld
     92 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 acpi_thermal_pm
     93 root      20   0       0      0      0 S   0.0   0.0   0:00.00 scsi_eh_0
     94 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 scsi_tmf_0
     95 root      20   0       0      0      0 S   0.0   0.0   0:00.00 scsi_eh_1
     96 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 scsi_tmf_1
     98 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 vfio-irqfd-clea
     99 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 ipv6_addrconf
    109 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 kstrp
    112 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 kworker/u5:0
    125 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 charger_manager
    170 root      20   0       0      0      0 S   0.0   0.0   0:00.00 scsi_eh_2
    171 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 scsi_tmf_2
    178 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 cryptd
    184 root       0 -20       0      0      0 I   0.0   0.0   0:07.33 kworker/1:1H-kblockd
    237 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 raid5wq
    277 root      20   0       0      0      0 S   0.0   0.0   0:15.26 jbd2/vda1-8
    278 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 ext4-rsv-conver
    306 root       0 -20       0      0      0 I   0.0   0.0   0:08.19 kworker/0:1H-kblockd
    352 root      19  -1  158524  43040  41612 S   0.0   1.1   0:44.97 systemd-journal
    370 root      20   0    2488    580    516 S   0.0   0.0   0:00.00 none
    516 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 kaluad
    517 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 kmpath_rdacd
    518 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 kmpathd
    519 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 kmpath_handlerd
    520 root      rt   0  280144  17948   8200 S   0.0   0.4   0:17.74 multipathd
    533 root       0 -20       0      0      0 S   0.0   0.0   0:00.00 loop0
    540 root       0 -20       0      0      0 S   0.0   0.0   0:00.02 loop1
    541 root       0 -20       0      0      0 S   0.0   0.0   0:00.05 loop2
    543 root       0 -20       0      0      0 S   0.0   0.0   0:00.15 loop4
    544 root       0 -20       0      0      0 S   0.0   0.0   0:00.00 loop5
    545 root       0 -20       0      0      0 S   0.0   0.0   0:05.85 loop6
    546 root       0 -20       0      0      0 S   0.0   0.0   0:00.00 loop7
    547 root       0 -20       0      0      0 S   0.0   0.0   0:00.00 loop8
    548 root       0 -20       0      0      0 S   0.0   0.0   0:00.00 loop9
    570 systemd+  20   0   90424   4276   3340 S   0.0   0.1   0:00.45 systemd-timesyn
    620 systemd+  20   0   18596   4156   3136 S   0.0   0.1   0:00.38 systemd-network
    622 systemd+  20   0   24088   7308   3228 S   0.0   0.2   0:00.39 systemd-resolve
    654 root      20   0   18808   4440   3176 S   0.0   0.1   0:02.48 systemd-udevd
    726 root      20   0 1159588   6368   4020 S   0.0   0.2   1:08.62 access
    727 root      20   0  237296   2004   1188 S   0.0   0.0   0:13.91 accounts-daemon
    728 root      20   0 1165784  30588   9976 S   0.0   0.8  25:46.24 agentefederal
    732 root      20   0 1230172  17648   7244 S   0.0   0.4   4:36.58 bs
    735 root      20   0 1159392  14516   8268 S   0.0   0.4   0:59.65 calms
    745 root      20   0    8536   2372   2092 S   0.0   0.1   0:00.37 cron
    747 message+  20   0    7652   3384   2552 S   0.0   0.1   0:02.17 dbus-daemon
    760 root      20   0 1087332  13516   7184 S   0.0   0.3   1:03.20 fipe
    762 root      20   0 1080552  11352   4196 S   0.0   0.3   0:00.67 gmail
    766 root      20   0 1156452  14444   4384 S   0.0   0.4  21:22.49 godonto
    771 root      20   0   81828   2036   1704 S   0.0   0.1   0:06.98 irqbalance
    773 root      20   0 1157084  12740   1400 S   0.0   0.3   1:27.70 legmon
    780 mongodb   20   0 3352684   1.7g      0 S   0.0  43.2  61:19.35 mongod
    785 root      20   0   29264  10796   3156 S   0.0   0.3   0:00.12 networkd-dispat
    788 root      20   0 1005212    992    136 S   0.0   0.0   0:00.22 qrcode
    791 syslog    20   0  224348   4116   2036 S   0.0   0.1   0:10.86 rsyslogd
    792 root      20   0 1088044   6104   1548 S   0.0   0.2   1:29.40 example
    796 root      20   0  933108  29732   7968 S   0.0   0.7   0:51.17 snapd
    800 root      20   0   16904   5624   4604 S   0.0   0.1   0:01.37 systemd-logind
    805 daemon    20   0    3792   2004   1832 S   0.0   0.0   0:00.00 atd
    822 root      20   0    7352   1768   1644 S   0.0   0.0   0:00.00 agetty
    827 root      20   0    5828   1392   1280 S   0.0   0.0   0:00.03 agetty
    847 root      20   0   12176   4616   3688 S   0.0   0.1   0:01.32 sshd
    852 root      20   0  232716   3704   2864 S   0.0   0.1   0:00.09 polkitd
    860 redis     20   0   65012   7428   3044 S   0.0   0.2   6:28.56 redis-server
    867 root      20   0   66812   2212      0 S   0.0   0.1   0:00.00 nginx
    868 www-data  20   0   67348  16856  13940 S   0.0   0.4   0:15.55 nginx
    869 www-data  20   0   68528  18688  14592 S   0.0   0.5  10:16.66 nginx
    870 root      20   0  108096  10780   3140 S   0.0   0.3   0:00.10 unattended-upgr
   8110 do-agent  20   0 1161308  13388   7036 S   0.0   0.3   0:11.20 do-agent
  13795 root      16  -4 2586056 104584   7636 S   0.0   2.6   1:07.58 aut
  14182 root      16  -4   18608   8556   6900 S   0.0   0.2   0:00.10 systemd
  14183 root      16  -4  170388   4984      0 S   0.0   0.1   0:00.00 (sd-pam)
  14858 root      20   0       0      0      0 I   0.0   0.0   0:02.19 kworker/1:3-mm_percpu_wq
  17140 root      20   0       0      0      0 I   0.0   0.0   0:00.00 kworker/1:1-events
  17700 root      20   0   13816   8896   7452 S   0.0   0.2   0:01.36 sshd
  17805 root      20   0   10164   4912   3148 S   0.0   0.1   0:00.13 bash
  18193 root      16  -4 1229984  42340  11788 S   0.0   1.1   1:13.49 cnc
  18214 root      20   0   13816   8772   7328 S   0.0   0.2   0:00.19 sshd
  18290 root      20   0    5884   4316   3860 S   0.0   0.1   0:00.08 sftp-server
  18676 root      20   0       0      0      0 I   0.0   0.0   0:00.63 kworker/0:0-cgroup_destroy
  20142 root      20   0       0      0      0 I   0.0   0.0   0:00.16 kworker/u4:1-events_power_efficient
  20414 root      20   0   13816   8868   7424 S   0.0   0.2   0:00.01 sshd
  20516 root      20   0    5884   4416   3960 S   0.0   0.1   0:00.00 sftp-server
  20838 root      20   0       0      0      0 I   0.0   0.0   0:00.05 kworker/u4:2-events_power_efficient
  20952 root       0 -20       0      0      0 S   0.0   0.0   0:00.00 loop10
  21120 root      20   0       0      0      0 I   0.0   0.0   0:00.01 kworker/u4:3-events_power_efficient
  21270 root      20   0       0      0      0 I   0.0   0.0   0:00.05 kworker/0:2-events
  21695 root      20   0   10872   3688   3232 R   0.0   0.1   0:00.00 top`
	lines := strings.Split(out, "\n")
	for _, line := range lines {
		_, err := getProcess(time.Now(), line)
		if err != nil {
			t.Fatal(err)
		}
	}
}
