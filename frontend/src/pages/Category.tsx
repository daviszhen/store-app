import { useEffect, useState } from 'react';
import { useParams, Link } from 'react-router-dom';
import type { Product } from '../types';
import { api } from '../api/client';

export default function Category() {
  const { id } = useParams<{ id: string }>();
  const [products, setProducts] = useState<Product[]>([]);

  useEffect(() => {
    if (id) api.getProducts(Number(id)).then(setProducts);
  }, [id]);

  return (
    <div className="page">
      <div className="page-title">商品列表</div>
      <div className="product-grid">
        {products.map((p) => (
          <Link key={p.id} to={`/product/${p.id}`} className="product-card">
            <div className="product-image">📦</div>
            <div className="product-info">
              <div className="product-name">{p.name}</div>
              <div className="product-price">¥{p.price}</div>
            </div>
          </Link>
        ))}
        {products.length === 0 && <div className="empty">暂无商品</div>}
      </div>
    </div>
  );
}
