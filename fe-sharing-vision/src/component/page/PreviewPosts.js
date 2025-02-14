import { useState, useEffect } from 'react';
import axios from 'axios';

export default function PreviewPosts() {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    const getPosts = async () => {
      try {
        const response = await axios.get(
          `http://localhost:8080/articles/publish/10/0`
        );
        setPosts(response.data);
      } catch (error) {
        console.error('Error fetching posts:', error);
      }
    };

    getPosts();
  }, []);

  return (
    <div className='max-w-3xl mx-auto p-5'>
      <h1 className='text-3xl font-bold mb-5'>Blog Posts</h1>
      {posts?.length > 0 ? (
        posts.map((post) => (
          <div key={post.id} className='bg-white p-4 rounded-lg shadow mb-4'>
            <h2 className='text-xl font-semibold'>{post.title}</h2>
            <p className='text-gray-600'>{post.content.substring(0, 150)}...</p>
            <span className='text-blue-500 text-sm'>{post.category}</span>
          </div>
        ))
      ) : (
        <p>No posts available</p>
      )}

    </div>
  );
}
