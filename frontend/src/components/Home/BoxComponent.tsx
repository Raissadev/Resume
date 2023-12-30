import { Row, Col, Layout, Button, Popover, Divider } from 'antd';
import { QqOutlined, RadarChartOutlined, HeatMapOutlined } from '@ant-design/icons';

function BoxComponent(): any
{
    return(
        <>
            <Layout className="boxes-component default">
                <Row justify="space-around">
                    <Col
                        span={6} xxl={6} xl={6} lg={6} md={24} sm={24} xs={24}
                    >
                        <Popover
                            content="Linux user, I've used distros like: Linux mint (only for testing on virtual machines rsrs), Ubuntu and currently Kali Linux and Manjaro"
                            title="Operational system"
                            trigger="hover"
                        >
                            <Button>
                                <QqOutlined />
                            </Button>
                        </Popover>
                        <Divider />
                    </Col>
                    <Col
                        span={6} xxl={6} xl={6} lg={6} md={24} sm={24} xs={24}
                    >
                        <Popover
                            content="I've worked with different technologies. In addition, the language that I introduced myself to the world of
                                computing was PHP, but I expanded my arsenal by adding stacks such as: JavaScript/TypeScript, Node, C, python, SQL and Golang...
                                I also like to use technologies aimed at devops like Docker and Terraform.
                            "
                            title="Technologies"
                            trigger="hover"
                        >
                            <Button>
                                <RadarChartOutlined />
                            </Button>
                        </Popover>
                        <Divider />
                    </Col>
                    <Col
                        span={6} xxl={6} xl={6} lg={6} md={24} sm={24} xs={24}
                    >
                        <Popover
                            content="I like calculus, rubik's cubes and anime. 🙂"
                            title="Brief introduction about me"
                            trigger="hover"
                        >
                            <Button>
                                <HeatMapOutlined />
                            </Button>
                        </Popover>
                        <Divider />
                    </Col>
                </Row>
            </Layout>
        </>
    );
}

export default BoxComponent;