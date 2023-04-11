type MailProperty =
{
    email: string,
    subject: string,
    message: string,
    response: string,
}

const MailPattern: MailProperty =
{
    email: "",
    subject: "",
    message: "",
    response: "",
}

export { type MailProperty, MailPattern };