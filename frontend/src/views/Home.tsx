import React, { useState } from "react";
import { FormProperty, FormPattern } from "../@types/form-modal";
import { Layout } from 'antd';
import BannerComponent from '../components/Home/BannerComponent';
import BoxComponent from '../components/Home/BoxComponent';
import InterestsComponent from '../components/Home/InterestsComponent';
import MessageComponent from '../components/Home/MessageComponent';
import Modall from "../components/Modal";

function Home(): any
{
    const [form, setForm] = useState<FormProperty>(FormPattern);

    return(
        <>
            <Layout>
                <BannerComponent />
            </Layout>

            <Layout>
                <MessageComponent />
            </Layout>

            <Layout>
                <BoxComponent />
            </Layout>

            <Layout>
                <InterestsComponent />
            </Layout>


            <Modall
                show={form.visibleModal}
                handleAction={() => setForm({...form, visibleModal: !form.visibleModal})}
            />
        </>
    );
}

export default Home;