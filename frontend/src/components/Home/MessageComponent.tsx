import { Layout, Typography, } from 'antd';

const { Title } = Typography;

function MessageComponent(): any
{
    return(
        <>
            <Layout className="message-component" id="work">
                <Title level={1}>
                    Hey! I'm Raissa, an 19 year old <br />
                    <span>software engineer</span>. I am passionate <br />
                    about the world of <span>computing</span>. Also, <br /> I like the
                    idea of transforming <br /> the world through <span>technology</span>.
                </Title>
            </Layout>
        </>
    );
}

export default MessageComponent;