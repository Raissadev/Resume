import { Row, Col, Layout, Typography, Divider, Image } from 'antd';
import networks from '../../assets/aot.png';
import back from '../../assets/ll.jpg';
import front from '../../assets/1344008.jpeg';

const { Title } = Typography;

function InterestsComponent(): any
{
    return(
        <>
            <Layout className="interests-component default">
                <Row>
                    <Col span={24}>
                        <Divider orientation="left" orientationMargin="0">
                            <Title level={2}>
                                DevOps
                            </Title>
                        </Divider>
                    </Col>
                    <Col>
                        <Image width={1000} src={networks} alt="Imagem meramente ilustrativa" />
                    </Col>
                </Row>
                <Row>
                    <Col span={24}>
                        <Divider orientation="right" orientationMargin="0">
                            <Title level={2}>
                                Backend
                            </Title>
                        </Divider>
                    </Col>
                    <Col>
                        <Image width={1000} src={back} alt="Imagem meramente ilustrativa" />
                    </Col>
                </Row>
                <Row>
                    <Col span={24}>
                        <Divider orientation="left" orientationMargin="0">
                            <Title level={2}>
                                Frontend
                            </Title>
                        </Divider>
                    </Col>
                    <Col>
                        <Image width={1000} src={front} alt="Imagem meramente ilustrativa" />
                    </Col>
                </Row>
            </Layout>
        </>
    );
}

export default InterestsComponent;