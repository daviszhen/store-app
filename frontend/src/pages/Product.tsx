import { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import type { Product } from '../types';
import { api } from '../api/client';
import { useCart } from '../context/CartContext';

export default function ProductPage() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const { addItem } = useCart();
  const [product, setProduct] = useState<Product | null>(null);
  const [quantity, setQuantity] = useState(1);
  const [adding, setAdding] = useState(false);

  useEffect(() => {
    if (id) api.getProduct(Number(id)).then(setProduct);
  }, [id]);

  const handleAdd = async () => {
    setAdding(true);
    await addItem(Number(id), quantity);
    setAdding(false);
    navigate('/cart');
  };

  if (!product) return <div className="loading">加载中...</div>;

  return (
    <div className="page">
      <div className="product-detail">
        <div className="pd-image">📦</div>
        <div className="pd-name">{product.name}</div>
        <div className="pd-price">¥{product.price}</div>
        {product.description && <div className="pd-desc">{product.description}</div>}

        <div className="pd-quantity">
          <span>数量</span>
          <div className="qty-control">
            <button onClick={() => setQuantity(Math.max(1, quantity - 1))}>-</button>
            <span>{quantity}</span>
            <button onClick={() => setQuantity(quantity + 1)}>+</button>
          </div>
        </div>

        <button className="btn-primary" onClick={handleAdd} disabled={adding}>
          {adding ? '添加中...' : '加入购物车'}
        </button>
      </div>
    </div>
  );
}
