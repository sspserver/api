+++
subject="Password Reset"
urlPasswordResetFormat="https://example.com/auth/reset-password?code={{code}}&email={{email}}"
+++

Dear {{.user.Email}},

We received a request to reset the password for your account associated with this email address. If you did not request a password reset, please ignore this email.

## Reset Your Password

To reset your password, please click on the link below:

[Reset Password]({{.urlPasswordReset}}) {{.urlPasswordReset}}

This link will expire in 1 hour. If you do not reset your password within this timeframe, you will need to submit a new request.

### Important Notes

- Do not share this email or the reset link with anyone.
- If you encounter any issues or did not request this change, please contact our support team immediately at {{.org.SupportEmail}}.

Thank you for using our services!

Best Regards,  
{{.org.Name}} Support Team
