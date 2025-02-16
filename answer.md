## Technical Questions

1. **Timezone Conflicts**: How would you handle timezone conflicts between participants
in an appointment?
= I would use the timezone of the user who created the appointment as the reference timezone. When a user creates an appointment, the user's timezone is stored in the database. When a user views the appointment, the appointment time is converted to the user's timezone. This ensures that the appointment time is displayed correctly for each user, regardless of their timezone.

2. **Database Optimization**: How can you optimize database queries to efficiently fetch
user-specific appointments?
= I would use indexes to optimize database queries. Indexes can speed up the retrieval of data by creating a data structure that allows the database to quickly locate the rows that match a query. I would create indexes on the columns that are frequently used in queries, such as the user_id column in the appointments table. This would allow the database to quickly find the appointments that belong to a specific user without having to scan the entire table.

3. **Additional Features**: If this application were to become a real product, what
additional features would you implement? Why?
= If this application were to become a real product, I would implement the following additional features:
   - Notifications: I would add a notification system that sends reminders to users about upcoming appointments. This would help users remember their appointments and reduce the likelihood of missed appointments.
   - Availability Calendar: I would add an availability calendar that shows users when they are available for appointments. This would make it easier for users to schedule appointments with each other by showing their availability at a glance
   - Integration with Calendar Apps: I would add the ability for users to sync their appointments with external calendar apps, such as Google Calendar or Outlook. This would allow users to view their appointments in their preferred calendar app and receive notifications about upcoming appointments.
  
4. **Session Management**: How would you manage user sessions securely while keeping
them lightweight (e.g., avoiding large JWT payloads)?
= I would use a jwt token to manage user sessions securely while keeping them lightweight. I would store only the user_id and expiration time in the jwt payload, keeping the payload small. I would also use short expiration times for the jwt tokens to reduce the risk of unauthorized access if a token is compromised. To handle session expiration, for next improvement, I would implement a refresh token mechanism to generate new jwt tokens without requiring the user to log in again.