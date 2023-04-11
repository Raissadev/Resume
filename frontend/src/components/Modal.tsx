import React, { useState, useEffect } from "react";
import api from '../services/api';
import { MailProperty, MailPattern } from "../@types/mail";
import { Modal, Layout, Form, Input, notification, Button } from 'antd';

const { TextArea } = Input;

function Modall({show, handleAction}: any): any
{
    const [loading, setLoading] = useState(false);
    const [form] = Form.useForm();

    useEffect((): any => {}, [ handleAction ]);

    const handleOk = async () => {
        form.validateFields().then(async (values) => {
            setLoading(true);
            await api.post("/mail", mail)
            .then((response: any) => {
                setLoading(false);
                notification.open({
                    message: "Email successfully sent",
                  });
            })
            .catch( (err: any) => {
                console.log('Invalid!');
            });
        }).catch();
    };
  
    const [mail, setMail]: any = useState<MailProperty>(MailPattern);

    return(
        <>
            <Layout>
                <Modal
                    title="Send Mail"
                    visible={show}
                    onOk={handleOk}
                    onCancel={handleAction}
                    footer={[
                        <Button key="back" onClick={handleAction}>
                          Cancel
                        </Button>,
                        <Button key="submit" type="primary" loading={loading} onClick={handleOk}>
                          Submit
                        </Button>,
                    ]}
                >
                    <Form
                        name="basic"
                        labelCol={{ span: 8 }}
                        wrapperCol={{ span: 16 }}
                        initialValues={{ remember: true }}
                        onFinish={handleOk}
                        form={form}
                        autoComplete="off"
                    >
                        <Form.Item
                            name="name"
                            rules={[{ required: true, message: 'Please input your name!' }]}
                        >
                            <Input
                                placeholder="Subject"
                                type="text"
                                onChange={(val: any) => setMail({ ...mail, subject: val.target.value })}
                            />
                        </Form.Item>
                        <Form.Item
                            name="email"
                            rules={[{ type: 'email', required: true, message: 'Please input your email!' }]}
                        >
                            <Input
                                placeholder="Email"
                                type="email"
                                onChange={(val: any) => setMail({ ...mail, email: val.target.value })}
                            />
                        </Form.Item>
                        <Form.Item
                            name="message"
                            rules={[{ required: true, message: 'Please input your message!' }]}
                        >
                            <TextArea
                                placeholder="Message"
                                onChange={(val: any) => setMail({ ...mail, message: val.target.value })}
                                rows={4}
                            />
                        </Form.Item>
                    </Form>
                </Modal>
            </Layout>
        </>
    );
}

export default Modall;