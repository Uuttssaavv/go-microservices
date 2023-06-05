package services

func (s *service) DeleteProduct(productId uint,userId uint) int {

	return s.repositories.DeleteProduct(productId,userId)
}
