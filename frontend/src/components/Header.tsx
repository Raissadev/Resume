import React, { useState } from "react";
import { FormProperty, FormPattern } from "../@types/form-modal";
import { Col, Row, Layout, Menu, Typography } from 'antd';
import Modall from "../components/Modal";

const { Header } = Layout;
const { Title } = Typography;

function Head(): any
{
    const [form, setForm] = useState<FormProperty>(FormPattern);

    return(
        <>
            <Layout>
                <Header>
                    <Row align="middle" justify="space-between">
                        <Col>
                            <Title level={4}>Raissadev</Title>
                        </Col>
                        <Col span={6} offset={6}>
                            <Menu mode="horizontal">
                                <Menu.Item key="home">
                                    <a href="#home">Home</a>
                                </Menu.Item>
                                <Menu.Item key="work">
                                    <a href="#work">Work</a>
                                </Menu.Item>
                                <Menu.Item key="media">
                                    <a href="#medias">Medias</a>
                                </Menu.Item>
                                <Menu.Item key="email" onClick={() => setForm({...form, visibleModal: !form.visibleModal})}>
                                    Email
                                </Menu.Item>
                            </Menu>
                        </Col>
                    </Row>
                </Header>

                <Modall
                    show={form.visibleModal}
                    handleAction={() => setForm({...form, visibleModal: !form.visibleModal})}
                />
            </Layout>
        </>
    );
}

export default Head;