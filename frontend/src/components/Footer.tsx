import React, { useState } from "react";
import { FormProperty, FormPattern } from "../@types/form-modal";
import { Row, Col, Layout, Typography, Button } from 'antd';
import { RightOutlined } from '@ant-design/icons';
import Modall from "../components/Modal";

const { Title } = Typography;
const { Footer } = Layout;

function Foot(): any
{
    const [form, setForm] = useState<FormProperty>(FormPattern);

    return(
        <>
            <Footer>
                <Row align="middle" justify="center">
                    <Col span={24}>
                        <Title level={2}>
                            Want to create <br/> amazing stuff
                        </Title>
                    </Col>
                    <Col span={24}>
                        <Button
                            type="primary"
                            size="large"
                            onClick={() => setForm({...form, visibleModal: !form.visibleModal})}
                        >
                            Send mail
                            <RightOutlined />
                        </Button>
                    </Col>
                </Row>
                <Row align="middle" justify="center" id="medias">
                    <Button href="https://github.com/Raissadev" type="link">GitHub</Button>
                    <Button href="https://www.linkedin.com/in/raissadev" type="link">Linkedin</Button>
                    <Button href="https://www.instagram.com/raissa_dev" type="link">Instagram</Button>
                    <Button href="https://twitter.com/Raissa_Dev" type="link">Twitter</Button>
                    <Button href="https://aur.archlinux.org/packages/kenbunshoku-haki" type="link">Kenbunshoku Haki Arch</Button>
                </Row>
                <Row align="middle" justify="center" className="copyright">
                    <Title level={4}>
                        ©2022 raissadev - All rights reserved
                    </Title>
                </Row>

                <Modall
                    show={form.visibleModal}
                    handleAction={() => setForm({...form, visibleModal: !form.visibleModal})}
                />
            </Footer>
        </>
    );
}

export default Foot;