# modak-rate-limiter

`Hey there fellow coder!` I sure hope this project lives up to the expectations, it basically:

 - Defines the quota for each type of notification in a certain timespan
 - Sends 15 notifications to 2 users
 - The rate limiter comes in, checking if the quota for the notification type is ok
 - If the notification is sent, its date is stored on the rate limiter for checking later quotas
 - The rate limiter also deletes old dates from the store when it doesn't need them anymore

`Have a nice day!`