# tender
🏗️ Tender Management System

Welcome to the Tender Management System, a comprehensive platform designed to simplify and enhance the tendering process for clients and contractors. Whether you're submitting bids, managing projects, or tracking notifications, our system ensures efficiency, transparency, and scalability.

✨ Features

For Clients:
Create Tenders: Post tenders with detailed descriptions, deadlines, and budgets.
Manage Tenders: Update tender statuses (e.g., Open, Closed, Awarded).
Seamless File Management: Attach project-related files securely.
For Contractors:
Submit Bids: Propose competitive bids with price, delivery time, and comments.
Bid Notifications: Stay updated with real-time notifications for tender activity.
User-Friendly Limits: Built-in safeguards ensure fair participation.
System Highlights:
Real-Time Validation: Prevents invalid operations like bidding on closed tenders.
Role-Based Access: Distinct roles for clients and contractors ensure secure and focused interactions.
Audit Trail: Timestamps and soft deletion for comprehensive data tracking.
🛠️ Tech Stack

Backend: PostgreSQL, PL/pgSQL (triggers, custom functions).
Framework: Built for scalability and modularity.
Languages: SQL, PL/pgSQL, with integrations ready for future extensions.
🚀 Installation Guide

Clone the repository:
git clone https://github.com/your-username/tender-management-system.git
cd tender-management-system
Set up the database:
Ensure PostgreSQL is installed.
Run the SQL schema provided in the db folder:
psql -U localhost -d tender -f schema.sql
Start the application:
Connect your backend to the database.
Configure your API server (if applicable).
📜 Database Structure

Tables
users: Stores user credentials and roles.
tenders: Manages tender details like deadlines, budgets, and statuses.
bids: Tracks contractor bids and associated information.
notifications: Handles user alerts for bids and tenders.
Triggers and Constraints
Prevents:
Bid creation for non-open tenders.
Status changes once a tender is marked as "awarded."
Limits contractors to 5 bids per minute.
🔒 Security

Role-based authorization.
Prevents invalid actions with robust triggers and database constraints.
📅 Roadmap

 Add support for tender categories.
 Introduce analytics for clients and contractors.
 Multi-language support for global reach.
 Integration with payment gateways for project milestones.
🤝 Contributing

The project uses a Makefile to simplify common tasks. Below is a breakdown of each command:

run_db: Starts the database and Redis containers in the background using Docker Compose.
stop_db: Stops the database and Redis containers.
run: Builds and runs the application containers using Docker Compose.
stop: Stops and removes all Docker containers associated with the project.
migrate: Runs database migrations.
logs: Tails logs of the running containers.
migrate_up: Applies all pending database migrations.
migrate_down: Rolls back the last batch of migrations.
migrate_force: Forces the database to a specific migration version.
migrate_file: Creates a new migration file for creating tables.
insert_file: Creates a new migration file for inserting data into tables.
build: Builds the application binary for macOS without CGo dependencies.
swag-gen: Generates Swagger API documentation.

We welcome contributions from developers around the globe! Please fork the repository, create a branch, and submit a pull request. For major changes, open an issue first to discuss your ideas.

🧑‍💻 Authors

This project was built with ❤️ by Feruza Mirjalilova, Sardorbek Axadjonov and Farrux Turdiqulov.

📄 License

This project is licensed under the MIT License. See the LICENSE file for details.

📢 Feedback

Got suggestions or need help? Open an issue or reach out via email at axadjonovsardorbeck@gmail.com.