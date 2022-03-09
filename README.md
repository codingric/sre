# SRE Challenge

## Q1

_A service daemon in production has stopped responding to network requests. You
receive an alert about the health of the service. Where would you start with handling
the alerting condition? How would you gather more information about the process
and its doing? What are common reasons a process might appear to be locked up,
and how would you rule out each possibility?_

Start at looking at metrics at wider view:
 - What was happening at the time of the alerts?
 - Are other related systems down?
 - Is the system saturated or at capacity?

Look closer the effected system:
 - Are other systems this depends on having issues?
 - Platform not scaling fast enough or at all?
 - 3rd party API/services having issues?

Look deeper into the system:
 - What do the logs show at the time of the alerts
 - Is there a pattern in errors logs?
 - Can you trace requests through the system to find the issue?

Analysis:
 - Use data collected to make a call of root cause
 - Idenity the effect on users
 - Formulate a work around or fix?
 - Implement work around or fix
 - Ensure resolved
 - Book Postmortem

Common reasons:
 - State data unavailable: Disc, configuration, database, API
 - Resources: RAM, CPU, Disc, nodes
 - Saturation: traffic/requests hit perfomance limits of the system therefore reducing throughput

---

## Q2

_You have a Linux binary you need to run called blackbox . When you attempt to run it
in a terminal (e.g. ./blackbox ), it prints a blank line and exits. What are the common
reasons a binary could behave this way? How would you troubleshoot it?_

 - Binary does output anything. eg. Only logs to file or is missing parameters?
    - Try running with -v flag for verbosity
    - Try running with -h to show options on how to use
    - man blackbox to see Manual?
 - Could be daemon? eg. Forks to create children processes?
    - Read documentation on running in the foreground if required?
 - Is it the right archeture for the system: eg. running Linux x86 binary on MacOS arm64 system
 - Are there missing libraries to mrun the binary? missing .dll .so .a .la

## Q3

_Write a program, topX - given a number X and an arbitrarily large file that contains
individual numbers on each line (e.g. 500Gb file), will output the largest X numbers,
highest first._


_Tell us about the run time/space complexity of it, and whether you think there's room
for improvement in your approach._

 - This application keeps track of numbers in memory in struct, this could potentially cause memory issues
 - Sort is built in quicksort and make O(n log n) comparisons
 - Could use a second file to keep track of numbers but could potentially double disc usuage
 - Tested with a 1GB file with 270M rows, completed in 19sec


_Deliverable for this is a GitHub repo with fully-working code in your preferred
language of choice._

[topX/main.go](topX/main.go)

## Q4

_One of the specifics of working in hospitality is a certain "seasonality" to our app.
Peaks in traffic correspond to Breakfast, Lunch, Dinner, increasing on
Thursday/Friday nights and pretty much all day Saturday. At the same time, we're
working in many timezones so they don't all correspond.

The application needs to scale to handle the load as each peak/trough window for a
timezone is observed. How would you design for this? Where do you think the main
bottlenecks could be? What actions would you take to understand the problem and
to finally deliver an infrastructure that would support it?_


 - An deep understand of all the moving parts in the system would be a start
 - Collecting and analysing metrics to understand staturation points
 - Delving into each components performance profiles would identify weak links in the chain where effort could be focused to strengthen and increase its ability to scale or deal with load
 - Design components to be more resilient to failure, fail gracefully and recover
 - Start from best bang for buck and work your way through until all risk is acceptable
 - Use of managed auto scaling functionality:
    - Cloud providers various mechanisms for scheduling scaling, AWS provides [ASG Scheduled actions](https://docs.aws.amazon.com/autoscaling/ec2/userguide/schedule_time.html#create-sch-actions)
    - Pods could be scaled with HPA and could potenially use [custom metrics](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#scaling-on-custom-metrics) that could force scaling up on predefined schedule
 - Auto scaling could still increase scaling but scheduled scaling would minimise ramp up times for new nodes or pods
 - Generally databases and external API are bottle necks to high load situations as scaling up events or rate limits could exacerbate any already fragile situation
 - Unfortuantly this question is _way_ too high level to give more specific answer
