import Navbar from "../components/Navbar";
import ProductGrid from "../components/ProductGrid";
import { Button, Space, Typography } from "antd";
import { Layout } from 'antd';


const { Title } = Typography;

const Home = () => {
  return (
    <Layout>
      <Layout style={{ background: '#141414', flex: 1 , minHeight: '100vh'}}>
      <div style={{ padding: '10px' }}>  {/* padding คือ ปรับขนาดของหน้า Structure */}
        <Navbar />
        <div style={{ background: 'linear-gradient(90deg, #9254de 0%, #f759ab 100%)', height: 180, borderRadius: 10, marginBottom: 24 }}></div> {/* borderRadius คือ ปรับมุมเหลี่ยม */} {/* code บรรทัดนี้ทำกรอบสี่เหลี่ยม */}
        <Title level={3} style={{ color: 'white' }}>Product</Title>
        <Space style={{ marginBottom: 16 }}>
          <Button type="primary" shape="round">Top</Button>
          <Button shape="round">Popular</Button>
          <Button shape="round">Recommended</Button>
          <Button shape="round">Filter</Button>
        </Space>
        <ProductGrid />
      </div>
      </Layout>
    </Layout>
  );
};

export default Home;
