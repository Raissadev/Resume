import React, { useState } from "react";
import { Row, Col, Layout, Image, Typography, Button } from 'antd';
import { RightOutlined, CaretLeftFilled } from '@ant-design/icons';
import { FormProperty, FormPattern } from "../../@types/form-modal";
import myImage from '../../assets/my.jpg';
import Modall from "../../components/Modal";

const { Title } = Typography;

function BannerComponent(): any
{
    const [form, setForm] = useState<FormProperty>(FormPattern);

    return(
        <>
            <Layout className="banner-home" id="home">
                <Row>
                    <Image width={320} src={myImage} />
                </Row>
                <Row>
                    <Col span={24}>
                        <Title level={1}>
                            INDEPENDENT FULL-STACK <br /> SOFTWARE ENGINNER
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
                <Row className="attach">
                    <Title level={4}>
                        <span>--------------- <CaretLeftFilled /></span>
                        SCROLL TO SEE MORE
                    </Title>
                </Row>

                <Modall
                    show={form.visibleModal}
                    handleAction={() => setForm({...form, visibleModal: !form.visibleModal})}
                />
            </Layout>
        </>
    );
}

export default BannerComponent;