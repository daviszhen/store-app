import { useState } from 'react';
import { useCart } from '../context/CartContext';
import { api } from '../api/client';

export default function Cart() {
  const { items, loading, total, count, updateQuantity, removeItem, refresh } = useCart();
  const [showForm, setShowForm] = useState(false);
  const [contact, setContact] = useState({ name: '', phone: '', addr: '' });
  const [remark, setRemark] = useState('');
  const [ordering, setOrdering] = useState(false);
  const [orderResult, setOrderResult] = useState('');

  const handleOrder = async () => {
    if (!contact.name || !contact.phone || !contact.addr) {
      alert('请填写完整联系方式');
      return;
    }
    setOrdering(true);
    try {
      const order: any = await api.createOrder(contact, remark);
      setOrderResult(`下单成功！订单号：${order.order_no}，金额：¥${order.total_amount}`);
      await refresh();
    } catch (e: any) {
      alert(e.message);
    } finally {
      setOrdering(false);
    }
  };

  if (loading) return <div className="loading">加载中...</div>;

  return (
    <div className="page">
      <div className="page-title">购物车 ({count})</div>

      {orderResult && <div className="order-success">{orderResult}</div>}

      {items.length === 0 && !orderResult && <div className="empty">购物车为空</div>}

      <div className="cart-list">
        {items.map((item) => (
          <div key={item.id} className="cart-item">
            <div className="ci-image">📦</div>
            <div className="ci-info">
              <div className="ci-name">{item.product?.name}</div>
              <div className="ci-price">¥{item.product?.price}</div>
            </div>
            <div className="ci-quantity">
              <button onClick={() => updateQuantity(item.id, item.quantity - 1)}>-</button>
              <span>{item.quantity}</span>
              <button onClick={() => updateQuantity(item.id, item.quantity + 1)}>+</button>
            </div>
            <button className="ci-delete" onClick={() => removeItem(item.id)}>删除</button>
          </div>
        ))}
      </div>

      {items.length > 0 && (
        <>
          <div className="cart-total">
            <span>合计：</span>
            <span className="total-price">¥{total.toFixed(2)}</span>
          </div>

          {!showForm ? (
            <button className="btn-primary" onClick={() => setShowForm(true)}>
              去下单
            </button>
          ) : (
            <div className="order-form">
              <input
                placeholder="收货人姓名"
                value={contact.name}
                onChange={(e) => setContact({ ...contact, name: e.target.value })}
              />
              <input
                placeholder="手机号码"
                value={contact.phone}
                onChange={(e) => setContact({ ...contact, phone: e.target.value })}
              />
              <input
                placeholder="收货地址"
                value={contact.addr}
                onChange={(e) => setContact({ ...contact, addr: e.target.value })}
              />
              <input
                placeholder="备注（选填）"
                value={remark}
                onChange={(e) => setRemark(e.target.value)}
              />
              <button className="btn-primary" onClick={handleOrder} disabled={ordering}>
                {ordering ? '提交中...' : '确认下单'}
              </button>
            </div>
          )}
        </>
      )}
    </div>
  );
}
