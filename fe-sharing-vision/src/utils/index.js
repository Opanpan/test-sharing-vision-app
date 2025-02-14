export function validateInputForm(data) {
  if (data.title.length < 19) {
    return 'Title min 20 Character';
  } else if (data.content.length < 199) {
    return 'Content min 200 Character';
  } else if (data.category.length < 4) {
    return 'Category min 3 Character';
  }
  return '';
}
