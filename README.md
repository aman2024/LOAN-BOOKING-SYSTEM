# LOAN-BOOKING-SYSTEM

Assumptions:
1. Authorization contains userId/adminId
2. Authorization is of type NoAuth with only value userId/adminId. Skipping encoding/decoding part
3. View api is returning all result of users (Not doing in pagination manner)
4. Admin blindly approves the loan without seeing his past history/balance
5. Not storing any term repayment history/backlog (since this was not mentioned in the question statement)


Steps to run the service
1. Install are the pre-requisite written in requirement.txt file
2. Set MYSQL username, host, port, password in .env.test file
3. Create table in MYSQL. DDL cmd in shared below
4. Run either of the cmds on terminal to run the service 
    a. ENV=dev go run .
    b. sh start.sh



MYSQL table Schema

CREATE TABLE `loan_info` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` varchar(45) NOT NULL,
  `amount` int NOT NULL,
  `term` int NOT NULL,
  `status` varchar(45) NOT NULL,
  `admin_approver_id` varchar(45) DEFAULT NULL,
  `term_paid` int DEFAULT '0',
  `amount_paid` int DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
)

We are Using single table to with name `loan_info`
`id`: This Is Unique AUTO_INCREMENT which refers to loan's id
`user_id` : This refers to users' id
`amount`: This refers to total amount user asked for loan
`term` : This refers to total term user asked for loan (it is numeric value)
`status`: This refers to current status of the loan. Can be in 3 states (PENDING,APPROVED,PAID)
`admin_approver_id`: This refers to id of admin who approved the loan
`term_paid`: This refers to total no of terms user paid till now
`amount_paid`: This refers to total amount paid by user in all his terms



