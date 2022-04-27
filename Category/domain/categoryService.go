package domain

type CategoryService interface {
	AddCategory(CategoryDescription) (string, error)
	GetCategories() ([]Category, error)
	GetCategoryById(string) (Category, error)
	UpdateCategoryById(Category) (bool, error)
	DeleteCategoryById(string) (bool, error)
	DeleteCategories([]string) (bool, error)
}

type categoryService struct {
	CategoryDynamoRepository CategoryDynamoRepository
}

func (service categoryService) AddCategory(categoryDescription CategoryDescription) (string, error) {
	id := GenerateUniqueId()
	categoryDescriptions := []CategoryDescription{}
	categoryDescriptions = append(categoryDescriptions, categoryDescription)
	categoryRecord := Category{
		Id:                  id,
		CategoryDescription: categoryDescriptions,
	}
	if ok, err := service.CategoryDynamoRepository.InsertCategory(categoryRecord); !ok {
		return id, err
	}
	return id, nil
}
func (service categoryService) GetCategories() ([]Category, error) {
	categoryRecords, err := service.CategoryDynamoRepository.FindAllCategories()
	if err != nil {
		return nil, err
	}
	return categoryRecords, nil
}

func (service categoryService) GetCategoryById(id string) (Category, error) {
	categoryRecord, err := service.CategoryDynamoRepository.FindCategoryByID(id)
	if err != nil {
		return Category{}, err
	}
	return categoryRecord, nil
}

func (service categoryService) UpdateCategoryById(category Category) (bool, error) {
	_, err := service.CategoryDynamoRepository.UpdateCategoryById(category)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (service categoryService) DeleteCategoryById(categoryId string) (bool, error) {
	_, err := service.CategoryDynamoRepository.DeleteCategoryById(categoryId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (service categoryService) DeleteCategories(categories []string) (bool, error) {
	_, err := service.CategoryDynamoRepository.DeleteCategories(categories)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewCategoryService(categoryDynamoRepository CategoryDynamoRepository) CategoryService {
	return &categoryService{
		CategoryDynamoRepository: categoryDynamoRepository,
	}
}
